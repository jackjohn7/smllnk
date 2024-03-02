package main

import (
	"fmt"
	"log"

	"github.com/jackjohn7/smllnk/app"
	"github.com/jackjohn7/smllnk/controllers"
	"github.com/jackjohn7/smllnk/db/repositories"
	"github.com/jackjohn7/smllnk/environment"
	mids "github.com/jackjohn7/smllnk/middlewares"
	"github.com/jackjohn7/smllnk/sessions"
)

func main() {
	// user sql repository (doesn't necessarily have to be postgres)
	repos := repositories.NewPGRepositories()
	// session storage (doesn't necessarily have to be redis)
	sessionStore, err := sessions.NewRedisSessionStore()
	if err != nil {
		log.Fatalln(err)
	}

	auth := mids.NewAuth("smllnk_session", repos, sessionStore)

	addr := fmt.Sprintf(":%s", environment.Env.Port)

	app := app.New(addr, []app.Controller{
		controllers.NewGeneralController(repos, sessionStore, auth),
		controllers.NewAccountsController(repos, sessionStore, auth),
	})
	// Handler will concat the paths here
	app.WithStaticDirectory("GET /public/styles/*", ".")
	app.WithStaticDirectory("GET /public/scripts/*", ".")

	// global middleware
	app.WithMiddleware(mids.NewLogger(app.Mux()).Start())

	// register controllers and serve
	app.Serve()
}
