package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/csrf"
	repos "github.com/jackjohn7/smllnk/db/repositories"
	mids "github.com/jackjohn7/smllnk/middlewares"
	"github.com/jackjohn7/smllnk/public/views/layout"
	"github.com/jackjohn7/smllnk/public/views/login"
	"github.com/jackjohn7/smllnk/sessions"
	"github.com/jackjohn7/smllnk/utils"
)

type AccountsController struct {
	repositories *repos.Repositories
	sessionStore sessions.SessionStore
	auth         *mids.Auth
}

func NewAccountsController(
	repo *repos.Repositories,
	sess sessions.SessionStore,
	auth *mids.Auth,
) *AccountsController {
	return &AccountsController{
		repositories: repo,
		sessionStore: sess,
		auth:         auth,
	}
}

func (c *AccountsController) Register(mux *http.ServeMux) error {
	mux.HandleFunc("GET /login", c.auth.AuthCtx(c.auth.RedirectIfAuthed("/", c.loginPageHandler)))
	mux.HandleFunc("POST /login", c.auth.AuthCtx(c.auth.RedirectIfAuthed("/", c.loginHandler)))
	mux.HandleFunc("GET /magic/{id}", c.auth.AuthCtx(c.auth.RedirectIfAuthed("/", c.magicHandler)))
	mux.HandleFunc("POST /logout", c.auth.AuthCtx(c.auth.Restrict(c.logoutHandler)))
	return nil
}

func (c *AccountsController) loginPageHandler(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, login.LoginPage(layout.BaseProps{
		Title:       "SmlLnk - Login",
		Description: "Get started sharing links today!",
		BaseUrl:     r.Host,
		AuthCtx:     nil,
		CsrfToken:   csrf.Token(r),
	}))
}

func (c *AccountsController) loginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.FormValue("email")
	if email == "" {
		// no email provided. Return Error
		w.WriteHeader(http.StatusOK)
		utils.Render(w, login.LoginTemplate(layout.BaseProps{
			Title:       "SmlLnk - Login",
			Description: "Get started sharing links today!",
			BaseUrl:     r.Host,
			AuthCtx:     nil,
			CsrfToken:   csrf.Token(r),
		}, "No email provided, lil bro"))
		return
	}

	// if there is no user, create one
	user, err := c.repositories.Users.GetByEmail(email)
	if err != nil {
		user, err = c.repositories.Users.Create(email)
		if err != nil {
			// if something goes wrong creating user, just write err (temp)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}

	// create magic request
	mr, err := c.repositories.MagicRequests.Create(user.Id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong creating Magic Login Link"))
		return
	}

	// temporarily just output mr.Id to stdout
	fmt.Printf("link: /magic/%s\n", mr.Id)

	utils.Render(w, login.CheckInbox(email))
}

func (c *AccountsController) magicHandler(w http.ResponseWriter, r *http.Request) {
	// get the session
	magicId := r.PathValue("id")

	mr, err := c.repositories.MagicRequests.Get(magicId)
	if err != nil {
		w.WriteHeader(http.StatusMovedPermanently)
		w.Write([]byte("Login link invalid"))
		return
	}

	// get user
	user, err := c.repositories.Users.GetById(mr.UserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("No user found. Perhaps you deleted your account?"))
		return
	}

	// create session
	session, err := c.sessionStore.Create(user, r.UserAgent())
	if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(err.Error()))
		return
	}

	// set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     c.auth.SessionCookieKey,
		Value:    session.Id,
		Expires:  session.ExpiresAt,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		HttpOnly: true,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (c *AccountsController) logoutHandler(w http.ResponseWriter, r *http.Request) {
	// get auth info
	acRaw := r.Context().Value("AuthCtx")
	if acRaw == nil {
		return
	}

	ac := acRaw.(*mids.AuthCtx)

	// delete session
	if ok := c.sessionStore.Delete(ac.Session.Id); !ok {
		fmt.Println("Something went wrong deleting session")
	}

	// set cookie to delete
	http.SetCookie(w, &http.Cookie{
		Name:     c.auth.SessionCookieKey,
		Value:    ac.Session.Id,
		Expires:  time.Now(),
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		HttpOnly: true,
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
