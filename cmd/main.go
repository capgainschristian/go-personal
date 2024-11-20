package main

import (
	"log"
	"net/http"

	"github.com/capgainschristian/go-personal/internal/pkg/render"
)

func main() {

	// Load and cache templates

	// Load routes
	render.RenderPages()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "home.page.tmpl", nil)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Println(err)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "about.page.tmpl", nil)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Println(err)
		}
	})

	log.Println("Starting server on :8080..")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}