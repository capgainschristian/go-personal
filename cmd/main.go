package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	// Load and cache templates

	// Load routes
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	templates := template.Must(template.ParseGlob("templates/*.tmpl"))

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