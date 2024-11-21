package main

import (
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

func main() {

	initialize()

	router := SetupRouter()

	log.Println("Starting server on :8080..")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func initialize() {
	templates = template.Must(template.ParseGlob("templates/*.tmpl"))

	for _, tmpl := range templates.Templates() {
        log.Println("Loaded template:", tmpl.Name())
    }
	
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}