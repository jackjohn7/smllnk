package controllers

import (
	"fmt"
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
	mux.HandleFunc("GET /login", loginPageHandler)
	mux.HandleFunc("POST /login", loginHandler)
	return nil
}

func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(200)
	// w.Write([]byte("Success"))
	utils.Render(w, login.LoginTemplate(csrf.Token(r), ""))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.FormValue("email")
	if email == "" {
		// no email provided. Return Error
		w.WriteHeader(http.StatusBadRequest)
		utils.Render(w, login.LoginTemplate(csrf.Token(r), "No email provided, lil bro"))
		return
	}

	// if there is no user, create one

	fmt.Println(email)
}
