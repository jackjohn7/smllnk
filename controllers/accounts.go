package controllers

import (
	"net/http"

	repos "github.com/jackjohn7/smllnk/db/repositories"
	mids "github.com/jackjohn7/smllnk/middlewares"
	"github.com/jackjohn7/smllnk/sessions"
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
	mux.HandleFunc("GET /login", loginHandler)
	return nil
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Success"))
}
