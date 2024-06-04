package controllers

import (
	"app/components"
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
