package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

type Controller interface {
	Register(*echo.Echo) error
}

type App struct {
	addr        string
	server      *echo.Echo
	controllers []Controller
}

func NewApp(addr string, controllers []Controller) *App {
	return &App{
		addr:        addr,
		server:      echo.New(),
		controllers: controllers,
	}
}

func (app *App) Server() *echo.Echo {
	return app.server
}

func (app *App) Serve() {
	// register controllers
	for _, cont := range app.controllers {
		cont.Register(app.Server())
	}

	// start with graceful shutdown stuff (from Echo cookbook)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
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
