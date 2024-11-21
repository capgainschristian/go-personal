package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var templates *template.Template

func main() {

	templates = parseTemplate()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	router := SetupRouter()

	log.Println("Starting server on :8080..")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func parseTemplate() *template.Template {
	template := template.New("")
	err := filepath.Walk("./templates", func(path string, info os.FileInfo, err error) error {
		log.Println("Processing file:", path)
		if strings.Contains(path, ".tmpl") {
			_, err := template.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}
		return err
	})
	if err != nil {
		panic(err)
	}

	for _, tmpl := range template.Templates() {
		log.Println("Loaded template:", tmpl.Name())
	}
	return template
}
