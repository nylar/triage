package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nylar/triage/models"
)

type ticketService struct {
	apiService
}

func FetchTickets(db *sql.DB) (*models.Tickets, error) {
	tkts := new(models.Tickets)

	rows, err := db.Query(`SELECT * FROM ticket`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tkt := new(models.Ticket)
		err := rows.Scan(&tkt.TicketID, &tkt.Message, &tkt.Status.StatusID)
		if err != nil {
			return nil, err
		}

		*tkts = append(*tkts, *tkt)
	}

	return tkts, nil
}

func (ts *ticketService) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	tickets, err := FetchTickets(ts.db)
	if err != nil {
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	enc.Encode(tickets)
}
