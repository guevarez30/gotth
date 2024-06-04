package api

import (
	"app/api/controllers"
	"app/api/middleware"
	"net/http"

	"app/views"

	"github.com/a-h/templ"
)

var redirectToLogin = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("exampleCookie")
	if err != nil {
		component := views.Layout(views.Index(), false)
		handler := templ.Handler(component)
		handler.ServeHTTP(w, r)
	} else {
		component := views.Layout(views.App(), true)
		handler := templ.Handler(component)
		handler.ServeHTTP(w, r)
	}
})

func Routes(mux *http.ServeMux) {

	mux.Handle("GET /", redirectToLogin)

	mux.Handle("POST /login", controllers.Login)

	mux.Handle("GET /login", controllers.LoginForm)

	mux.Handle("POST /logout", controllers.Logout)

	// Auth Required
	mux.Handle("POST /ping", middleware.Chain(controllers.Ping, middleware.AuthRequired))

	mux.Handle("GET /nav-items", middleware.Chain(controllers.NavItems, middleware.AuthRequired))

	mux.Handle("GET /content", middleware.Chain(controllers.Ping, middleware.AuthRequired))

	mux.Handle("GET /close/popup", middleware.Chain(controllers.ClosePopup, middleware.AuthRequired))
}
