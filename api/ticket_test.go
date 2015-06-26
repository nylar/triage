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
	r, err := http.NewRequest("GET", "/api/tickets/", nil)
	if err != nil {
		t.Error(err.Error())
	}

	w := httptest.NewRecorder()
	Routes(router, conn)
	router.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code)
}

func TestTicketIndexError(t *testing.T) {
	db, err := sqlmock.New()
	if err != nil {
		t.Errorf(err.Error())
	}

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

	if err = db.Close(); err != nil {
		t.Errorf("Error '%s' was not expected while closing the database", err)
	}

}

func TestFetchTickets(t *testing.T) {
	_, err := FetchTickets(conn)
	assert.NoError(t, err)
}

func TestFetchTicketsRowsError(t *testing.T) {
	db, err := sqlmock.New()
	if err != nil {
		t.Errorf(err.Error())
	}

	sqlmock.ExpectQuery("SELECT (.+) FROM ticket").
		WillReturnError(fmt.Errorf("Query failed"))

	_, err = FetchTickets(db)
	assert.Error(t, err)

	if err = db.Close(); err != nil {
		t.Errorf("Error '%s' was not expected while closing the database", err)
	}
}

func TestFetchTicketsScanError(t *testing.T) {
	db, err := sqlmock.New()
	if err != nil {
		t.Errorf(err.Error())
	}

	sqlmock.ExpectQuery("SELECT (.+) FROM ticket").
		WillReturnRows(sqlmock.
		NewRows([]string{"ticket_id", "message"}).
		AddRow("x", "y"))

	_, err = FetchTickets(db)
	assert.Error(t, err)

	if err = db.Close(); err != nil {
		t.Errorf("Error '%s' was not expected while closing the database", err)
	}
}
