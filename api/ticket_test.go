package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllTickets(t *testing.T) {
	r, err := http.NewRequest("GET", "/api/tickets", nil)
	if err != nil {
		t.Error(err.Error())
	}

	w := httptest.NewRecorder()
	Routes(router)
	router.ServeHTTP(w, r)

	assert.Equal(t, 501, w.Code)
}
