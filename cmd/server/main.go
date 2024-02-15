package main

import (
	"github.com/jackjohn7/smllnk/app"
	"github.com/jackjohn7/smllnk/controllers"
	mids "github.com/jackjohn7/smllnk/middlewares"
)

func main() {
	app := app.NewApp(":3005", []app.Controller{
		controllers.NewGeneralController(),
	})

	// global middleware
	app.Server().Use(mids.NewLogger(app.Server()).Start())

	// register controllers and serve
	app.Serve()
}
