package api // import "github.com/nylar/triage/api/v1"

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/nylar/triage"
	"github.com/sirupsen/logrus"
)

type TicketService struct {
	db *sqlx.DB
}

func (ts *TicketService) View() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			logrus.WithError(err).Errorln("Couldn't parse ID parameter")
			http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
			return
		}

		ticket := &triage.Ticket{}
		if err := ticket.FindByID(ts.db, id); err != nil {
			logrus.WithError(err).Errorln("Couldn't fetch ticket")
			switch {
			case err == sql.ErrNoRows:
				http.Error(w, "Ticket not found", http.StatusNotFound)
			default:
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ticket)
	})
}

func (ts *TicketService) List() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tickets := &triage.Tickets{}
		if err := tickets.FindAll(ts.db); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tickets)
	})
}
