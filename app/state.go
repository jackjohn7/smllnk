package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

/*
A controller registers related handler functions and assigns them to endpoints

A controller must have the `Register` method defined.
*/
type Controller interface {
	Register(*echo.Echo) error
}

/*
App is what contains the echo server, registered controllers, and global state.
*/
type App struct {
	addr        string
	server      *echo.Echo
	controllers []Controller
}

/*
Creates a new App struct with an address and a set of controllers
*/
func New(addr string, controllers []Controller) *App {
	return &App{
		addr:        addr,
		server:      echo.New(),
		controllers: controllers,
	}
}

/*
Supplies middleware to echo server
*/
func (app *App) WithMiddleware(f echo.MiddlewareFunc) *App {
	app.server.Use(f)
	return app
}

/*
Returns the Echo server of the App.

This is useful for adding global middleware
*/
func (app *App) Server() *echo.Echo {
	return app.server
}

/*
Registers controllers and starts Echo server with graceful shutdown
*/
func (app *App) Serve() {
	// register controllers
	for _, cont := range app.controllers {
		cont.Register(app.Server())
	}

	// start with graceful shutdown stuff (from Echo cookbook)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server in separate goroutine
	go func() {
		if err := app.server.Start(app.addr); err != nil {
			log.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
