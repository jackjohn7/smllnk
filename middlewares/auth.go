package middlewares

import (
	"net/http"
	"time"

	"github.com/jackjohn7/smllnk/db/models"
	"github.com/jackjohn7/smllnk/db/repositories"
	"github.com/jackjohn7/smllnk/sessions"
	"github.com/labstack/echo/v4"
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

func (a *Auth) AuthCtx() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// validate that the user has a session
			// get session token from cookies
			cookie, err := c.Cookie(a.sessionCookieKey)
			if err != nil {
				return next(c)
			}

			sessionId := cookie.Value
			session, err := a.store.Get(sessionId)
			if err != nil {
				// their session seems to be erroneous or expired. Redirect to login and clear cookie
				c.SetCookie(&http.Cookie{
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
			c.Set("AuthCtx", authCtx)

			return next(c)
		}
	}
}

func (a *Auth) Restrict() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// get auth info
			ac := c.Get("AuthCtx")
			if ac == nil {
				return c.Redirect(http.StatusSeeOther, "/login")
			}
			return next(c)
		}
	}
}
