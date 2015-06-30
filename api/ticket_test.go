package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestTicketIndex(t *testing.T) {
	db := setUp()
	defer tearDown(db)

	sqlmock.ExpectQuery("SELECT (.+) FROM ticket").
		WillReturnRows(sqlmock.
		NewRows([]string{"ticket_id", "message", "status_id"}).
		AddRow(1, "y", 1))

	r, err := http.NewRequest("GET", "/api/tickets/", nil)
	if err != nil {
		t.Error(err.Error())
	}

	w := httptest.NewRecorder()
	h := TicketIndex(db)
	h.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code)

	Routes(router, db)
	router.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code)
}

func TestTicketIndexError(t *testing.T) {
	db := setUp()
	defer tearDown(db)

	sqlmock.ExpectQuery("SELECT (.+) FROM ticket").
		WillReturnError(fmt.Errorf("Query failed"))

	r, err := http.NewRequest("GET", "/api/tickets/", nil)
	if err != nil {
		t.Error(err.Error())
	}

	w := httptest.NewRecorder()
	h := TicketIndex(db)
	h.ServeHTTP(w, r)

	assert.Equal(t, 500, w.Code)
}

func TestFetchTickets(t *testing.T) {
	db := setUp()
	defer tearDown(db)

	sqlmock.ExpectQuery("SELECT (.+) FROM ticket").
		WillReturnRows(sqlmock.
		NewRows([]string{"ticket_id", "message", "status_id"}).
		AddRow(1, "y", 1))

	_, err := FetchTickets(db)
	assert.NoError(t, err)
}

func TestFetchTicketsRowsError(t *testing.T) {
	db := setUp()
	defer tearDown(db)

	sqlmock.ExpectQuery("SELECT (.+) FROM ticket").
		WillReturnError(fmt.Errorf("Query failed"))

	_, err := FetchTickets(db)
	assert.Error(t, err)
}

func TestFetchTicketsScanError(t *testing.T) {
	db := setUp()
	defer tearDown(db)

	sqlmock.ExpectQuery("SELECT (.+) FROM ticket").
		WillReturnRows(sqlmock.
		NewRows([]string{"ticket_id", "message", "status_id"}).
		AddRow("x", "y", "z"))

	_, err := FetchTickets(db)
	assert.Error(t, err)
}
