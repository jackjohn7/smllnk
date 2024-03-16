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
	// this route should be DELETE but DELETE was being forbidden for some reason
	mux.HandleFunc("POST /links/{linkId}", c.auth.AuthCtx(c.auth.Restrict(c.deleteLinkHandler)))
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

func (c *LinksController) deleteLinkHandler(w http.ResponseWriter, r *http.Request) {
	linkId := r.PathValue("linkId")
	if linkId == "" {
		// return err
	}
	acRaw := r.Context().Value("AuthCtx")
	if acRaw == nil {
		// if not authed, return 403
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You must log in to use this functionality"))
		return
	}

	// verify that user owns this particular link
	ac := acRaw.(*mids.AuthCtx)
	links, err := c.repositories.Links.GetAllUserLinks(ac.User)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// render the link with a message saying that it couldn't be deleted
		w.Write([]byte("")) // temp
		return
	}

	// I would like to write a function dedicated to this to allow SQL to
	//  perform the heavy lifting on this check, but this will work for the
	//  time being. This is temporary logic.

	// find link in links
	found := false
	for _, l := range links {
		if l.Id == linkId {
			found = true
			break
		}
	}

	if found {
		if ok := c.repositories.Links.Delete(linkId); ok {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(""))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			// render the link with a message saying that it couldn't be deleted
			w.Write([]byte("")) // temp
		}
	} else {
		// the user doesn't have access to this resource
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(""))
	}
}
