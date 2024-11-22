package main

import (
	"log"
	"net/http"

	"github.com/capgainschristian/go-personal/internal/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("home page")
	render.RenderPages(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	log.Println("about page")
	render.RenderPages(w, "about.page.tmpl")
}
