package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/nylar/triage/models"
)

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

func StatusIndex(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		enc := json.NewEncoder(w)

		statuses, err := FetchStatuses(db)
		if err != nil {
			errorResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}

		enc.Encode(statuses)
	})
}
