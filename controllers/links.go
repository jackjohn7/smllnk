package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/csrf"
	repos "github.com/jackjohn7/smllnk/db/repositories"
	mids "github.com/jackjohn7/smllnk/middlewares"
	"github.com/jackjohn7/smllnk/public/views/components"
	"github.com/jackjohn7/smllnk/public/views/layout"
	"github.com/jackjohn7/smllnk/sessions"
	"github.com/jackjohn7/smllnk/utils"
)

type LinksController struct {
	repositories *repos.Repositories
	sessionStore sessions.SessionStore
	auth         *mids.Auth
}

func NewLinksController(
	repo *repos.Repositories,
	sess sessions.SessionStore,
	auth *mids.Auth,
) *LinksController {
	return &LinksController{
		repositories: repo,
		sessionStore: sess,
		auth:         auth,
	}
}

func (c *LinksController) Register(mux *http.ServeMux) error {
	mux.HandleFunc("POST /links", c.auth.AuthCtx(c.auth.Restrict(c.createLinkHandler)))
	return nil
}

func (c *LinksController) createLinkHandler(w http.ResponseWriter, r *http.Request) {
	acRaw := r.Context().Value("AuthCtx")
	if acRaw == nil {
		// if not authed, return 403
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You must log in to use this functionality"))
		return
	}

	r.ParseForm()
	destination := r.FormValue("destination")
	if destination == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No destination provided"))
		return
	}
	name := r.FormValue("nickname")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No name provided"))
		return
	}

	ac := acRaw.(*mids.AuthCtx)

	// create link
	link, err := c.repositories.Links.Create(destination, name, ac.User)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("err: %s", err.Error())))
		return
	}

	baseProps := layout.BaseProps{
		AuthCtx:   ac,
		CsrfToken: csrf.Token(r),
	}
	utils.Render(w, components.Link(baseProps, components.LinkProps{
		Link: *link,
	}))
}
