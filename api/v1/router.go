package api // import "github.com/nylar/triage/api/v1"

import (
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

// Router sets up the API v1 routes
func Router(db *sqlx.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api/v1/", func(r chi.Router) {
		r.Route("/project", func(r chi.Router) {
			projectService := &ProjectService{
				db: db,
			}
			r.Get("/", projectService.List())
			r.Get("/{id}", projectService.View())
		})

	})

	return r
}
