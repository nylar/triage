package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type apiService struct {
	db *sql.DB
}

type errResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func Routes(r *httprouter.Router, db *sql.DB) {
	api := apiService{db: db}

	// Ticket API routes
	ticket := ticketService{api}
	r.POST("/api/tickets", ticket.Create)
	r.GET("/api/tickets/:id", ticket.Show)
	r.GET("/api/tickets", ticket.Index)

	// Status API routes
	status := statusService{api}
	r.GET("/api/statuses", status.Index)
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
