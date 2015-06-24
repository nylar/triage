package api

import "github.com/gorilla/mux"

type errResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type response struct {
	StatusCode int      `json:"status_code"`
	Data       struct{} `json:"data"`
}

func Routes(m *mux.Router) {
	s := m.PathPrefix("/api").Subrouter()

	s.Handle("/tickets", AllTickets()).Methods("GET")
}
