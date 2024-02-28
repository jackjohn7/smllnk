package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/urfave/negroni"
)

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

/*
A controller registers related handler functions and assigns them to endpoints

A controller must have the `Register` method defined.
*/
type Controller interface {
	Register(*http.ServeMux) error
}

type staticDirectory struct {
	urlPath    string
	fileServer http.Handler
}

/*
App is what contains the echo server, registered controllers, and global state.
*/
type App struct {
	addr        string
	mux         *http.ServeMux
	server      *http.Server
	middlewares []MiddlewareFunc
	controllers []Controller
	staticPaths []*staticDirectory
}

/*
Creates a new App struct with an address and a set of controllers
*/
func New(addr string, controllers []Controller) *App {
	return &App{
		addr:        addr,
		mux:         http.NewServeMux(),
		middlewares: make([]MiddlewareFunc, 0),
		controllers: controllers,
		staticPaths: make([]*staticDirectory, 0),
	}
}

/*
Supplies middleware to echo server
*/
func (app *App) WithMiddleware(f MiddlewareFunc) *App {
	app.middlewares = append(app.middlewares, f)
	return app
}

/*
Serve static directory
*/
func (app *App) WithStaticDirectory(urlPath string, filePath string) *App {
	app.staticPaths = append(app.staticPaths, &staticDirectory{
		urlPath:    urlPath,
		fileServer: http.FileServer(http.Dir(filePath)),
	})
	return app
}

/*
Returns the *http.Server of the App.
*/
func (app *App) Server() *http.Server {
	return app.server
}

/*
Return *http.ServeMux of the application
*/
func (app *App) Mux() *http.ServeMux {
	return app.mux
}

/*
Registers controllers and starts Echo server with graceful shutdown
*/
func (app *App) Serve() {
	// register controllers
	for _, cont := range app.controllers {
		cont.Register(app.mux)
	}

	for _, sd := range app.staticPaths {
		app.mux.Handle(sd.urlPath, sd.fileServer)
	}

	// Using negroni's panic-recovery middleware to prevent panics from destroying the app
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.UseHandler(app.mux)

	app.server = &http.Server{
		Addr:    app.addr,
		Handler: ApplyMiddleware(n, app.middlewares),
	}

	// start with graceful shutdown stuff (from Echo cookbook)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server in separate goroutine
	go func() {
		if err := app.server.ListenAndServe(); err != nil {
			log.Fatalf("shutting down the server: %s", err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.server.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}
