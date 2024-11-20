package render

import (
	"html/template"
	"net/http"
)

func RenderPages() *template.Template {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	templates := template.Must(template.ParseGlob("templates/*.tmpl"))
	
	return templates
}