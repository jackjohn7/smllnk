package main

import (
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"

	"github.com/jackjohn7/smllnk/app"
	"github.com/jackjohn7/smllnk/controllers"
	"github.com/jackjohn7/smllnk/db/repositories"
	"github.com/jackjohn7/smllnk/environment"
	mids "github.com/jackjohn7/smllnk/middlewares"
)

func main() {
	// user sql repository (doesn't necessarily have to be postgres)
	repos := repositories.NewPGRepositories()

	addr := fmt.Sprintf(":%s", environment.Env.Port)

	app := app.New(addr, []app.Controller{
		controllers.NewGeneralController(),
	}).WithRepositories(repos)

	// global middleware
	app.WithMiddleware(session.Middleware(
		sessions.NewCookieStore(
			environment.Env.AuthEnv.SessionSecret,
		),
	))
	app.WithMiddleware(mids.NewLogger(app.Server()).Start())

	// register controllers and serve
	app.Serve()
}
