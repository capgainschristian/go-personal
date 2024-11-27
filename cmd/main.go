package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	r := SetupRouter()

	log.Println("Starting server on :8080..")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
