package main

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"

	"github.com/jackjohn7/smllnk/app"
	"github.com/jackjohn7/smllnk/controllers"
	"github.com/jackjohn7/smllnk/db/repositories"
	mids "github.com/jackjohn7/smllnk/middlewares"
)

func main() {
	// user sql repository (doesn't necessarily have to be postgres)
	repos := repositories.NewPGRepositories(nil)

	app := app.New(":3005", []app.Controller{
		controllers.NewGeneralController(),
	}).WithRepositories(repos).
		WithMiddleware(session.Middleware(
			sessions.NewCookieStore([]byte("secret"))), // replace later
		)

	// global middleware
	app.WithMiddleware(mids.NewLogger(app.Server()).Start())

	// register controllers and serve
	app.Serve()
}
