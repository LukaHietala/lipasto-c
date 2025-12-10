package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			ImageURL string
		}{
			ImageURL: "https://images.dog.ceo/breeds/pomeranian/n02112018_10309.jpg",
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
			http.Error(w, "template rendering failed", http.StatusInternalServerError)
			return
		}
	})

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
