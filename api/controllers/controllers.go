package controllers

import (
	"net/http"

	"github.com/a-h/templ"
)

func Render(component templ.Component, w http.ResponseWriter, r *http.Request) {
	handler := templ.Handler(component)
	handler.ServeHTTP(w, r)
}
