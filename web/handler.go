package web

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
		r.Get("/{u_id}", h.FacultyList())
	})
	h.Route("/faculties", func(r chi.Router) {
		r.Get("/add", h.AddFaculty())
		r.Post("/add", h.AddFaculty())
		r.Get("/{f_id}", h.FacultyMemberList())
	})
	h.Route("/members", func(r chi.Router) {
		r.Get("/add", h.AddFacultyMember())
		r.Post("/add", h.AddFacultyMember())
	})

	h.Route("/api", func(r chi.Router) {
		r.Get("/universities/{u_id}/faculties", h.GetFaculties())
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
