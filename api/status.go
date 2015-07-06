package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nylar/triage/models"
)

type statusService struct {
	apiService
}

func FetchStatuses(db *sql.DB) (*models.Statuses, error) {
	statuses := new(models.Statuses)

	rows, err := db.Query(`SELECT * FROM status`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		status := new(models.Status)
		err := rows.Scan(&status.StatusID, &status.Name)
		if err != nil {
			return nil, err
		}

		*statuses = append(*statuses, *status)
	}

	return statuses, nil
}

func (ss *statusService) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")

	enc := json.NewEncoder(w)

	statuses, err := FetchStatuses(ss.db)
	if err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	enc.Encode(statuses)
}
