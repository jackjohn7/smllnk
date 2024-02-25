package controllers

import (
	"github.com/jackjohn7/smllnk/db/repositories"
	mids "github.com/jackjohn7/smllnk/middlewares"
	"github.com/jackjohn7/smllnk/sessions"

	"github.com/labstack/echo/v4"
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

func (c *GeneralController) Register(app *echo.Echo) error {
	app.GET("/", IndexHandler, c.auth.AuthCtx(), c.auth.Restrict())
	return nil
}

func IndexHandler(c echo.Context) error {
	return c.String(200, "Hello, world")
}
