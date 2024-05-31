package controllers

import (
	"app/templates/components"
	"app/templates/layouts"
	"app/views"
	"net/http"

	"github.com/a-h/templ"
)

func Render(component templ.Component, w http.ResponseWriter, r *http.Request) {
	handler := templ.Handler(component)
	handler.ServeHTTP(w, r)
}

var NavItems = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var navItems = []string{"banana", "apple"}
	component := components.NavItems(navItems)
	Render(component, w, r)
})

var Ping = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	component := components.Ping()
	Render(component, w, r)
})

var Login = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "exampleCookie",
		Value:    "Hello world!",
		Path:     "/",
		MaxAge:   600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)

	component := layouts.Layout(views.App(), true)
	Render(component, w, r)
})

var Logout = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "exampleCookie",
		Value:    "Hello world!",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	component := layouts.Layout(views.Index(), false)
	Render(component, w, r)
})
