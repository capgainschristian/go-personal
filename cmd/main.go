package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var templates *template.Template

func main() {

	templates = parseTemplates()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	router := SetupRouter()

	log.Println("Starting server on :8080..")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func parseTemplates() *template.Template {
	root := template.New("root") // Explicitly name the root template
	err := filepath.Walk("./templates", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // Handle any walking error
		}

		// Only process files with .tmpl extension
		if !info.IsDir() && strings.HasSuffix(path, ".tmpl") {
			log.Println("Processing file:", path)

			// ParseFiles returns the updated template; assign it back to root
			updatedRoot, parseErr := root.ParseFiles(path)
			if parseErr != nil {
				log.Printf("Error parsing file %s: %v", path, parseErr)
				return parseErr
			}
			root = updatedRoot // Update the root with the new template
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the templates directory: %v", err)
	}

	// Log all loaded templates
	for _, tmpl := range root.Templates() {
		log.Println("Loaded template:", tmpl.Name())
	}

	return root
}
