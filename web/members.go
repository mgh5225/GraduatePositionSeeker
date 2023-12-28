package web

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/mgh5225/gps"
)

func (h *Handler) FacultyMemberList() http.HandlerFunc {

	type data struct {
		Mems []gps.FacultyMember
	}

	files := layoutFiles()
	files = append([]string{"templates/members.html"}, files...)

	tmpl := template.Must(template.ParseFiles(files...))

	return func(w http.ResponseWriter, r *http.Request) {

		f_id, parse_err := uuid.Parse(chi.URLParam(r, "f_id"))

		if parse_err != nil {
			http.Error(w, parse_err.Error(), http.StatusInternalServerError)
			return
		}

		fm, err := h.store.FacultyMembers(f_id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.ExecuteTemplate(w, "bootstrap", data{Mems: fm})
	}
}

func (h *Handler) AddFacultyMember() http.HandlerFunc {

	type data struct {
		Univs []gps.University
	}

	files := layoutFiles()
	files = append([]string{"templates/add-faculty-member.html"}, files...)

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
			firstName := r.FormValue("first-name")
			lastName := r.FormValue("last-name")
			email := r.FormValue("email")
			f_idStr := r.FormValue("f-id")

			f_id, parse_err := uuid.Parse(f_idStr)

			if parse_err != nil {
				http.Error(w, parse_err.Error(), http.StatusInternalServerError)
				return
			}

			if err := h.store.CreateFacultyMember(&gps.FacultyMember{
				ID:        uuid.New(),
				FacultyID: f_id,
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
			}); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, fmt.Sprintf("/faculties/%s", f_idStr), http.StatusFound)

		}
	}
}
