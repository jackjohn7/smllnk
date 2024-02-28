package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jackjohn7/smllnk/db/models"
	"github.com/jackjohn7/smllnk/db/repositories"
	"github.com/jackjohn7/smllnk/sessions"
)

type (
	Auth struct {
		sessionCookieKey string
		repos            *repositories.Repositories
		store            sessions.SessionStore
	}

	AuthCtx struct {
		User    *models.User
		Session *sessions.Session
	}
)

func NewAuth(key string, sessStore sessions.SessionStore) *Auth {
	return &Auth{
		sessionCookieKey: key,
		store:            sessStore,
	}
}

func (a *Auth) AuthCtx(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// validate that the user has a session
		// get session token from cookies
		cookie, err := r.Cookie(a.sessionCookieKey)
		if err != nil {
			next(w, r)
			return
		}

		sessionId := cookie.Value
		session, err := a.store.Get(sessionId)
		if err != nil {
			// their session seems to be erroneous or expired. Redirect to login and clear cookie
			r.AddCookie(&http.Cookie{
				Name:     a.sessionCookieKey,
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
		fmt.Println("Here in auth")
		// get auth info
		ac := r.Context().Value("AuthCtx")
		if ac == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			next(w, r)
		}
	}
}
