package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() http.Handler {
	mux := chi.NewRouter()
	mux.Get("/", Home)
	mux.Get("/about", About)
	return mux
}