package controllers

import (
	repos "github.com/jackjohn7/smllnk/db/repositories"
	mids "github.com/jackjohn7/smllnk/middlewares"
	"github.com/jackjohn7/smllnk/sessions"

	"github.com/labstack/echo/v4"
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

func (c *AccountsController) Register(app *echo.Echo) error {
	app.GET("/login", loginHandler)
	return nil
}

func loginHandler(c echo.Context) error {
	return c.String(200, "Login Page")
}
