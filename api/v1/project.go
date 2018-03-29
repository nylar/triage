package api // import "github.com/nylar/triage/api/v1"

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/nylar/triage/triage"
	"github.com/sirupsen/logrus"
)

type ProjectService struct {
	db *sqlx.DB
}

func (ps *ProjectService) View() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			logrus.WithError(err).Errorln("Couldn't parse ID parameter")
			http.Error(w, "Invalid project ID", http.StatusBadRequest)
			return
		}

		project := &triage.Project{}
		if err := project.FindByID(ps.db, id); err != nil {
			logrus.WithError(err).Errorln("Couldn't fetch project")
			switch {
			case err == sql.ErrNoRows:
				http.Error(w, "Project not found", http.StatusNotFound)
			default:
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(project)
	})
}

func (ps *ProjectService) List() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		projects := &triage.Projects{}
		if err := projects.FindAll(ps.db); err != nil {
			logrus.WithError(err).Errorln("Couldn't fetch all projects")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(projects)
	})
}

func (ps *ProjectService) Tickets() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			logrus.WithError(err).Errorln("Couldn't parse ID parameter")
			http.Error(w, "Invalid project ID", http.StatusBadRequest)
			return
		}

		tickets := &triage.Tickets{}
		if err := tickets.FindByProjectID(ps.db, id); err != nil {
			logrus.WithError(err).WithField("project id", id).
				Errorln("Couldn't find tickets for project")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tickets)
	})
}
