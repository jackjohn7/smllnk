package controllers

import (
	"net/http"

	"github.com/gorilla/csrf"
	repos "github.com/jackjohn7/smllnk/db/repositories"
	mids "github.com/jackjohn7/smllnk/middlewares"
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
	return nil
}

func (c *AccountsController) loginPageHandler(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(200)
	// w.Write([]byte("Success"))
	utils.Render(w, login.LoginTemplate(csrf.Token(r), ""))
}

func (c *AccountsController) loginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.FormValue("email")
	if email == "" {
		// no email provided. Return Error
		w.WriteHeader(http.StatusBadRequest)
		utils.Render(w, login.LoginTemplate(csrf.Token(r), "No email provided, lil bro"))
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
		}
	}

	/*
		Temporarily, I'm just going to immediately log in the user.
		In the future, the user should be emailed a magic link instead
	*/

	// create session for user
	session, err := c.sessionStore.Create(user, r.UserAgent())
	if err != nil {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(err.Error()))
		return
	}

	// set session cookie (TEMP)
	http.SetCookie(w, &http.Cookie{
		Name:     c.auth.SessionCookieKey,
		Value:    session.Id,
		Expires:  session.ExpiresAt,
		SameSite: http.SameSiteStrictMode,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther) // in the future, redirect to value in query param
}
