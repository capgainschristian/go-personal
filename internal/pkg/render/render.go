package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var TemplateCache map[string]*template.Template

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// Retrieve the requested template from the cache
	t, ok := TemplateCache[tmpl]
	if !ok {
		log.Printf("Template %s not found in cache", tmpl)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Use buf to ensure atomic writes and better error checking
	buf := new(bytes.Buffer)
	err := t.Execute(buf, nil)
	if err != nil {
		log.Println("Error rendering template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Write the rendered template to the response writer
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing to response:", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//page.tmpl files have to be rendered first before layout.tmpl files.
	//That is why filepath.Glob being called twice.
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		//if there are layout files, parse and add them to ts
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
