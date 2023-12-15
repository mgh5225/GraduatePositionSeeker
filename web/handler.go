package web

import (
	"encoding/json"
	"html/template"
	"net/http"

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
	})

	return h
}

type Handler struct {
	*chi.Mux

	store gps.Store
}

func (h *Handler) Index() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func (h *Handler) UniversityList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		u, err := h.store.Universities()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(u)
	}
}
