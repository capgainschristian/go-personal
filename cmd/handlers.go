package main

import (
	"log"
	"net/http"
)

templates = render.RenderPages()

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "home.page.tmpl", nil)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
	}
}