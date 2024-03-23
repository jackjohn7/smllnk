package controllers

import (
	"net/http"
	"os"

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
	mux.HandleFunc("GET /", c.auth.AuthCtx(c.auth.Restrict(c.indexHandler)))
	mux.HandleFunc("GET /favicon.ico", c.faviconHandler)
	return nil
}

func (c *GeneralController) indexHandler(w http.ResponseWriter, r *http.Request) {
	acRaw := r.Context().Value("AuthCtx")
	if acRaw == nil {
		return
	}

	ac := acRaw.(*mids.AuthCtx)
	links, err := c.repositories.Links.GetAllUserLinks(ac.User)
	if err != nil {
		// render err template in future
	}
	utils.Render(w, index.IndexTemplate(layout.BaseProps{
		Title:       "SmlLnk",
		Description: "Simplest and Cheapest Link-shortener",
		BaseUrl:     r.Host,
		AuthCtx:     ac,
		CsrfToken:   csrf.Token(r),
	}, links))
}

func (c *GeneralController) faviconHandler(w http.ResponseWriter, r *http.Request) {
	buf, err := os.ReadFile("public/favicon.png")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("nah"))
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(buf)
}
