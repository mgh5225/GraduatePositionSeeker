package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/mgh5225/gps"
)

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

func (h *Handler) GetFaculties() http.HandlerFunc {

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

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(f)
	}
}
