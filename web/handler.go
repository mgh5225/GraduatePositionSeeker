package web

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/mgh5225/gps"
)

func NewHandler(store gps.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	fs := http.FileServer(http.Dir("static/"))

	h.Use(middleware.Logger)
	h.Get("/", h.Index())
	h.Handle("/static/*", http.StripPrefix("/static/", fs))
	h.Route("/universities", func(r chi.Router) {
		r.Get("/", h.UniversityList())
		r.Get("/add", h.AddUniversity())
		r.Post("/add", h.AddUniversity())
	})

	return h
}

func layoutFiles() []string {
	files, err := filepath.Glob("templates/layouts/*.html")
	if err != nil {
		panic(err)
	}
	return files
}

type Handler struct {
	*chi.Mux

	store gps.Store
}

func (h *Handler) Index() http.HandlerFunc {
	files := layoutFiles()
	files = append([]string{"templates/index.html"}, files...)

	tmpl := template.Must(template.ParseFiles(files...))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "bootstrap", nil)
	}
}

func (h *Handler) UniversityList() http.HandlerFunc {
	type data struct {
		Univs []gps.University
	}

	files := layoutFiles()
	files = append([]string{"templates/universities.html"}, files...)

	tmpl := template.Must(template.ParseFiles(files...))

	return func(w http.ResponseWriter, r *http.Request) {

		u, err := h.store.Universities()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.ExecuteTemplate(w, "bootstrap", data{Univs: u})
	}
}

func (h *Handler) AddUniversity() http.HandlerFunc {

	files := layoutFiles()
	files = append([]string{"templates/add-university.html"}, files...)

	tmpl := template.Must(template.ParseFiles(files...))

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			tmpl.ExecuteTemplate(w, "bootstrap", nil)
			return
		case http.MethodPost:
			name := r.FormValue("name")

			if err := h.store.CreateUniversity(&gps.University{
				ID:   uuid.New(),
				Name: name,
			}); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/universities", http.StatusFound)

		}
	}
}
