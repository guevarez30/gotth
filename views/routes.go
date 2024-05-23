package views

import (
	"app/api/controllers"
	"app/api/middleware"
	"app/templates/components"
	"app/templates/layouts"
	"net/http"

	"github.com/a-h/templ"
)

var loggedIn = false

// Render middleware calls templ.Handler with the provided component
func Render(component templ.Component) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler := templ.Handler(component)
		handler.ServeHTTP(w, r)
	})
}

var redirectToLogin = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// authHeader := r.Header.Get("Authorization")
	if loggedIn {
		component := layouts.Layout(App(), true)
		handler := templ.Handler(component)
		handler.ServeHTTP(w, r)
	} else {
		component := layouts.Layout(Index(), false)
		handler := templ.Handler(component)
		handler.ServeHTTP(w, r)
	}
})

func login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loggedIn = true
		component := layouts.Layout(App(), true)
		handler := templ.Handler(component)
		handler.ServeHTTP(w, r)
	})
}

func logout() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loggedIn = false
		component := layouts.Layout(Index(), false)
		handler := templ.Handler(component)
		handler.ServeHTTP(w, r)
	})
}

func Routes(mux *http.ServeMux) {

	mux.Handle("GET /", redirectToLogin)

	mux.Handle("POST /ping", Render(components.Ping()))

	mux.Handle("POST /login", login())

	mux.Handle("POST /logout", logout())

	mux.Handle("GET /nav-items", middleware.Chain(controllers.GetItems))

	mux.Handle("GET /content", middleware.Chain(Render(components.Ping())))
}
