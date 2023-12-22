package web

import (
	"fmt"
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
		r.Get("/{u_id}", h.FacultyList())
	})
	h.Route("/faculties", func(r chi.Router) {
		r.Get("/add", h.AddFaculty())
		r.Post("/add", h.AddFaculty())
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

func (h *Handler) FacultyList() http.HandlerFunc {

	type data struct {
		Facs []gps.Faculty
	}

	files := layoutFiles()
	files = append([]string{"templates/faculties.html"}, files...)

	tmpl := template.Must(template.ParseFiles(files...))

	return func(w http.ResponseWriter, r *http.Request) {

		u_id, parse_err := uuid.Parse(chi.URLParam(r, "u_id"))

		if parse_err != nil {
			http.Error(w, parse_err.Error(), http.StatusInternalServerError)
			return
		}

		f, err := h.store.Faculties(u_id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.ExecuteTemplate(w, "bootstrap", data{Facs: f})
	}
}

func (h *Handler) AddFaculty() http.HandlerFunc {

	type data struct {
		Univs []gps.University
	}

	files := layoutFiles()
	files = append([]string{"templates/add-faculty.html"}, files...)

	tmpl := template.Must(template.ParseFiles(files...))

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:

			u, err := h.store.Universities()

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			tmpl.ExecuteTemplate(w, "bootstrap", data{Univs: u})
			return
		case http.MethodPost:
			name := r.FormValue("name")
			u_idStr := r.FormValue("u-id")

			u_id, parse_err := uuid.Parse(u_idStr)

			if parse_err != nil {
				http.Error(w, parse_err.Error(), http.StatusInternalServerError)
				return
			}

			if err := h.store.CreateFaculty(&gps.Faculty{
				ID:           uuid.New(),
				UniversityID: u_id,
				Name:         name,
			}); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, fmt.Sprintf("/universities/%s", u_idStr), http.StatusFound)

		}
	}
}
