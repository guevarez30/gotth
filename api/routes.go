package api

import (
	"net/http"

	"app/views"

	"github.com/a-h/templ"
)

func Render(component templ.Component, w http.ResponseWriter, r *http.Request) {
	handler := templ.Handler(component)
	handler.ServeHTTP(w, r)
}

var ViewController = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	component := views.Layout()
	Render(component, w, r)
})

func Routes(mux *http.ServeMux) {
	mux.Handle("GET /", ViewController)

}
