package handlers

import (
	"html/template"
	"net/http"

	templates "github.com/AleksandrBorohovsky/photo_portfolio"
	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	tmpl := template.Must(template.ParseFS(templates.TemplatesFS,
		"web/templates/layouts/*.html",
		"web/templates/pages/*.html",
		"web/templates/partials/*.html",
	))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Title string
		}{
			Title: "home page",
		}

		err := tmpl.ExecuteTemplate(w, "base", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	return r
}
