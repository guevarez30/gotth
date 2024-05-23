package controllers

import (
	"app/templates/components"
	"net/http"

	"github.com/a-h/templ"
)

var GetItems = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var navItems = []string{"banana", "apple"}
	templ.Handler(components.NavItems(navItems)).ServeHTTP(w, r)
})
