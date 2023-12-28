package web

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"github.com/mgh5225/gps"
)

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
