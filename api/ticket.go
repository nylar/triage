package api

import (
	"encoding/json"
	"net/http"
)

func AllTickets() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotImplemented)

		enc := json.NewEncoder(w)

		enc.Encode(errResponse{
			StatusCode: 501,
			Message:    "Endpoint not yet implemented",
		})
	})
}
