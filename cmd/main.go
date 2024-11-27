package main

import (
	"log"
	"net/http"

	"github.com/capgainschristian/go-personal/internal/pkg/render"
)

func main() {

	var err error
	render.TemplateCache, err = render.CreateTemplateCache()
	if err != nil {
		log.Fatalf("Error creating template cache: %v", err)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	r := SetupRouter()

	log.Println("Starting server on :8080..")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
