package main

import (
	"log"
	"net/http"

	"github.com/capgainschristian/go-personal/internal/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("home page")
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	log.Println("about page")
	render.RenderTemplate(w, "about.page.tmpl")
}
