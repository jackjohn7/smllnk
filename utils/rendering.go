package utils

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
)

func Render(w http.ResponseWriter, component templ.Component) error {
	return component.Render(context.Background(), w)
}
