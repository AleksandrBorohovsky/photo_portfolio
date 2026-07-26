package handlers

import (
	"html/template"
	"log/slog"
	"net/http"
	"time"

	templates "github.com/AleksandrBorohovsky/photo_portfolio"
	"github.com/go-chi/chi"
)

func NewRouter(log *slog.Logger) *chi.Mux {
	r := chi.NewRouter()

	fileServer := http.FileServer(http.Dir("web/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	shared := template.Must(template.ParseFS(
		templates.TemplatesFS,
		"web/templates/layouts/*.html",
		"web/templates/partials/*.html",
	))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log = log.With(
			slog.String("op", "GET:/"),
		)

		tmpl := template.Must(shared.Clone())
		tmpl = template.Must(tmpl.ParseFS(
			templates.TemplatesFS,
			"web/templates/pages/index.html",
		))

		data := struct {
			Title string
			Year  int
		}{
			Title: "Портретный фотограф в Минске",
			Year:  time.Now().Year(),
		}

		err := tmpl.ExecuteTemplate(w, "base", data)
		if err != nil {
			log.Error("failed to execute template", slog.String("err", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Get("/individual", func(w http.ResponseWriter, r *http.Request) {
		log = log.With(
			slog.String("op", "GET:/individual"),
		)

		tmpl := template.Must(shared.Clone())
		tmpl = template.Must(tmpl.ParseFS(
			templates.TemplatesFS,
			"web/templates/pages/individual.html",
		))

		err := tmpl.ExecuteTemplate(w, "base", nil)
		if err != nil {
			log.Error("failed to execute template", slog.String("err", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Get("/love-story", func(w http.ResponseWriter, r *http.Request) {
		log = log.With(
			slog.String("op", "GET:/love-story"),
		)

		tmpl := template.Must(shared.Clone())
		tmpl = template.Must(tmpl.ParseFS(
			templates.TemplatesFS,
			"web/templates/pages/love-story.html",
		))

		err := tmpl.ExecuteTemplate(w, "base", nil)
		if err != nil {
			log.Error("failed to execute template", slog.String("err", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Get("/family", func(w http.ResponseWriter, r *http.Request) {
		log = log.With(
			slog.String("op", "GET:/family"),
		)

		tmpl := template.Must(shared.Clone())
		tmpl = template.Must(tmpl.ParseFS(
			templates.TemplatesFS,
			"web/templates/pages/family.html",
		))

		err := tmpl.ExecuteTemplate(w, "base", nil)
		if err != nil {
			log.Error("failed to execute template", slog.String("err", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Get("/content", func(w http.ResponseWriter, r *http.Request) {
		log = log.With(
			slog.String("op", "GET:/content"),
		)

		tmpl := template.Must(shared.Clone())
		tmpl = template.Must(tmpl.ParseFS(
			templates.TemplatesFS,
			"web/templates/pages/content.html",
		))

		err := tmpl.ExecuteTemplate(w, "base", nil)
		if err != nil {
			log.Error("failed to execute template", slog.String("err", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Get("/events", func(w http.ResponseWriter, r *http.Request) {
		log = log.With(
			slog.String("op", "GET:/events"),
		)

		tmpl := template.Must(shared.Clone())
		tmpl = template.Must(tmpl.ParseFS(
			templates.TemplatesFS,
			"web/templates/pages/events.html",
		))

		err := tmpl.ExecuteTemplate(w, "base", nil)
		if err != nil {
			log.Error("failed to execute template", slog.String("err", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Get("/contacts", func(w http.ResponseWriter, r *http.Request) {
		log = log.With(
			slog.String("op", "GET:/contacts"),
		)

		tmpl := template.Must(shared.Clone())
		tmpl = template.Must(tmpl.ParseFS(
			templates.TemplatesFS,
			"web/templates/pages/contacts.html",
		))

		err := tmpl.ExecuteTemplate(w, "base", nil)
		if err != nil {
			log.Error("failed to execute template", slog.String("err", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	return r
}
