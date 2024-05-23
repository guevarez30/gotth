package views

import (
	"app/templates/components"
	"app/templates/layouts"
	"net/http"

	"github.com/a-h/templ"
)

func Routes(mux *http.ServeMux) {
	mux.Handle("GET /", templ.Handler(layouts.Layout(Index())))

	mux.Handle("POST /ping", templ.Handler(components.Ping()))
}
