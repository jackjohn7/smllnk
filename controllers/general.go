package controllers

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/jackjohn7/smllnk/db/repositories"
	mids "github.com/jackjohn7/smllnk/middlewares"
	"github.com/jackjohn7/smllnk/public/views/index"
	"github.com/jackjohn7/smllnk/public/views/layout"
	"github.com/jackjohn7/smllnk/sessions"
	"github.com/jackjohn7/smllnk/utils"
)

type GeneralController struct {
	repositories *repositories.Repositories
	sessionStore sessions.SessionStore
	auth         *mids.Auth
}

func NewGeneralController(
	repo *repositories.Repositories,
	sess sessions.SessionStore,
	auth *mids.Auth,
) *GeneralController {
	return &GeneralController{
		repositories: repo,
		sessionStore: sess,
		auth:         auth,
	}
}

func (c *GeneralController) Register(mux *http.ServeMux) error {
	mux.HandleFunc("GET /", c.auth.AuthCtx(c.auth.Restrict(indexHandler)))
	return nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	acRaw := r.Context().Value("AuthCtx")
	if acRaw == nil {
		return
	}

	ac := acRaw.(*mids.AuthCtx)
	utils.Render(w, index.IndexTemplate(layout.BaseProps{
		Title: "SmlLnk",
		Description: "Simplest and Cheapest Link-shortener",
		AuthCtx: ac,
		CsrfToken: csrf.Token(r),
	}))
}
