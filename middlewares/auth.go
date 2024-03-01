package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/jackjohn7/smllnk/db/models"
	"github.com/jackjohn7/smllnk/db/repositories"
	"github.com/jackjohn7/smllnk/sessions"
)

type (
	Auth struct {
		SessionCookieKey string
		repos            *repositories.Repositories
		store            sessions.SessionStore
	}

	AuthCtx struct {
		User    *models.User
		Session *sessions.Session
	}
)

func NewAuth(key string, repos *repositories.Repositories, sessStore sessions.SessionStore) *Auth {
	return &Auth{
		SessionCookieKey: key,
		repos:            repos,
		store:            sessStore,
	}
}

func (a *Auth) AuthCtx(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// validate that the user has a session
		// get session token from cookies
		cookie, err := r.Cookie(a.SessionCookieKey)
		if err != nil {
			next(w, r)
			return
		}

		sessionId := cookie.Value
		session, err := a.store.Get(sessionId)
		if err != nil {
			// their session seems to be erroneous or expired. Redirect to login and clear cookie
			http.SetCookie(w, &http.Cookie{
				Name:     a.SessionCookieKey,
				Value:    "",
				Path:     "/",
				Expires:  time.Unix(0, 0),
				HttpOnly: true,
			})
		}
		// if we have session, go fetch corresponding user
		user, err := a.repos.Users.GetById(session.UserId)

		// extend context with auth and user info
		authCtx := &AuthCtx{User: user, Session: session}
		r = r.WithContext(context.WithValue(r.Context(), "AuthCtx", authCtx))

		next(w, r)
	}
}

func (a *Auth) Restrict(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get auth info
		ac := r.Context().Value("AuthCtx")
		if ac == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			next(w, r)
		}
	}
}

func (a *Auth) RedirectIfAuthed(destination string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get auth info
		ac := r.Context().Value("AuthCtx")
		if ac != nil {
			http.Redirect(w, r, destination, http.StatusSeeOther)
		} else {
			next(w, r)
		}
	}
}
