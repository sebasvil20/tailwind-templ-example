package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/sebasvil20/tailwind-templ-example/views"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.Handle("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}))

	http.Handle("/", templ.Handler(views.Index()))

	http.ListenAndServe(":3000", nil)
}
