package controllers

import (
	"app/components"
	"net/http"
)

var ClosePopup = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	component := components.Nothing()
	Render(component, w, r)
})
