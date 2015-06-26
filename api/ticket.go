package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/nylar/triage/models"
)

func FetchTickets(db *sql.DB) (*models.Tickets, error) {
	tkts := new(models.Tickets)

	rows, err := db.Query(`SELECT * FROM ticket`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tkt := new(models.Ticket)
		err := rows.Scan(&tkt.TicketID, &tkt.Message)
		if err != nil {
			return nil, err
		}

		*tkts = append(*tkts, *tkt)
	}

	return tkts, nil
}

func TicketIndex(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		enc := json.NewEncoder(w)

		tkts, err := FetchTickets(db)
		if err != nil {
			errorResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}

		enc.Encode(tkts)
	})
}
