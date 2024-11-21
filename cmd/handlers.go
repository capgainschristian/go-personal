package main

import (
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "home.page.tmpl", nil)
	log.Println("home page")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "about.page.tmpl", nil)
	log.Println("about page")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println(err)
	}
}
