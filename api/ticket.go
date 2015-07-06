package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nylar/triage/models"
	"github.com/nylar/triage/utils"
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

func FetchTicket(db *sql.DB, id int) (*models.Ticket, error) {
	tkt := new(models.Ticket)

	stmt, err := db.Prepare(`SELECT * FROM ticket WHERE ticket_id = $1`)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&tkt.TicketID, &tkt.Message, &tkt.Status.StatusID)
	if err != nil {
		return nil, err
	}

	return tkt, nil
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

func (ts *ticketService) Show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	id := p.ByName("id")
	tid, err := utils.ParseToInt(id)
	if err != nil {
		errorResponse(w, "ID param could not be parsed as an integer", http.StatusBadRequest)
		return
	}

	ticket, err := FetchTicket(ts.db, tid)
	switch {
	case err == sql.ErrNoRows:
		errorResponse(w, "Ticket not found", http.StatusNotFound)
		return
	case err != nil:
		errorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	enc.Encode(ticket)
}
