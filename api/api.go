package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type errResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func Routes(m *mux.Router, db *sql.DB) {
	s := m.PathPrefix("/api").Subrouter()

	s.Handle("/tickets/", TicketIndex(db)).Methods("GET")
	s.Handle("/statuses/", StatusIndex(db)).Methods("GET")
}

func errorResponse(w http.ResponseWriter, err string, code int) {
	enc := json.NewEncoder(w)
	w.WriteHeader(code)
	enc.Encode(errResponse{
		StatusCode: code,
		Message:    err,
	})
	return
}
