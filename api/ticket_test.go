package api

import (
	"fmt"
	"net/http"
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

	route := server.URL + "/api/tickets"
	resp, err := http.Get(route)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestTicketIndexError(t *testing.T) {
	db := setUp()
	defer tearDown(db)

	sqlmock.ExpectQuery("SELECT (.+) FROM ticket").
		WillReturnError(fmt.Errorf("Query failed"))

	route := server.URL + "/api/tickets"
	resp, err := http.Get(route)

	assert.NoError(t, err)
	assert.Equal(t, 500, resp.StatusCode)
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
