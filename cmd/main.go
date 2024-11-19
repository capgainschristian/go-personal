package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	templates := template.Must(template.ParseFiles(
		filepath.Join("templates", "base.layout.tmpl"),
		filepath.Join("templates", "home.page.tmpl"),
	))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "home.page.tmpl", nil)
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