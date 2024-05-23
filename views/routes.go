package views

import (
	NavController "app/api/controllers"
	"app/api/middleware"
	"app/templates/components"
	"app/templates/layouts"
	"net/http"

	"github.com/a-h/templ"
)

// Render middleware calls templ.Handler with the provided component
func Render(component templ.Component) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler := templ.Handler(component)
		handler.ServeHTTP(w, r)
	})
}

func Routes(mux *http.ServeMux) {

	mux.Handle("GET /", Render(layouts.Layout(Index())))

	mux.Handle("POST /ping", Render(components.Ping()))

	mux.Handle("GET /nav-items", middleware.Chain(NavController.GetItems))

	mux.Handle("GET /content", middleware.Chain(Render(components.Ping())))
}
