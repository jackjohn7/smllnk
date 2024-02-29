package app

import "net/http"

func applyMiddleware(h http.Handler, middlewares []MiddlewareFunc) http.Handler {
	for _, m := range middlewares {
		h = m(h.ServeHTTP)
	}
	return h
}
