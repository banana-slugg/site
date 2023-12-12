package main

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

const (
	BASE_TEMPLATE = "templates/base.html"
)

func main() {
	// boilerplate
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// index route, greeting page
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, r, "templates/pages/index.html")
	})
	// temp page at the moment
	r.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, r, "templates/pages/about.html")
	})
	// fileserver for css generated by tailwind
	fileServer := http.FileServer(http.Dir("./dist/"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	//turn that server on
	http.ListenAndServe(":3000", r)
}

// helper function to render the template for any page
func RenderTemplate(w http.ResponseWriter, r *http.Request, filename string) {
	template := template.Must(template.ParseFiles(filename, BASE_TEMPLATE))
	var buf bytes.Buffer
	if err := template.ExecuteTemplate(&buf, "base", nil); err != nil {
		//todo: throw up a 404 page here
		return
	}

	render.HTML(w, r, buf.String())
}
