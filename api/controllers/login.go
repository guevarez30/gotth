package controllers

import (
	"app/components"
	"app/views"
	"net/http"
)

var LoginForm = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	login := components.LoginForm()
	component := components.PopUp("Login", login)
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

	component := views.Layout(views.App(), true)
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
	component := views.Layout(views.Index(), false)
	Render(component, w, r)
})
