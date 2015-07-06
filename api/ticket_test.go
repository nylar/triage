package api

import (
	"database/sql"
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

func TestTicketShow(t *testing.T) {
	db := setUp()
	defer tearDown(db)

	sqlmock.ExpectPrepare()
	sqlmock.ExpectQuery("SELECT (.+) FROM ticket WHERE ticket_id = \\$1").
		WithArgs(1).
		WillReturnRows(sqlmock.
		NewRows([]string{"ticket_id", "message", "status_id"}).
		AddRow(1, "y", 1))

	route := server.URL + "/api/tickets/1"
	resp, err := http.Get(route)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestTicketShowUnparsableID(t *testing.T) {
	db := setUp()
	defer tearDown(db)

	route := server.URL + "/api/tickets/one"
	resp, err := http.Get(route)

	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestTicketShow404Error(t *testing.T) {
	db := setUp()
	defer tearDown(db)

	sqlmock.ExpectPrepare()
	sqlmock.ExpectQuery("SELECT (.+) FROM ticket WHERE ticket_id = \\$1").
		WithArgs(1).
		WillReturnError(sql.ErrNoRows)

	route := server.URL + "/api/tickets/1"
	resp, err := http.Get(route)

	assert.NoError(t, err)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestTicketShowError(t *testing.T) {
	db := setUp()
	defer tearDown(db)

	sqlmock.ExpectPrepare()
	sqlmock.ExpectQuery("SELECT (.+) FROM ticket WHERE ticket_id = \\$1").
		WithArgs(1).
		WillReturnError(fmt.Errorf("Some db error"))

	route := server.URL + "/api/tickets/1"
	resp, err := http.Get(route)

	assert.NoError(t, err)
	assert.Equal(t, 500, resp.StatusCode)
}

func TestFetchTicket(t *testing.T) {
	db := setUp()
	defer tearDown(db)

	sqlmock.ExpectPrepare()
	sqlmock.ExpectQuery("SELECT (.+) FROM ticket WHERE ticket_id = \\$1").
		WithArgs(1).
		WillReturnRows(sqlmock.
		NewRows([]string{"ticket_id", "message", "status_id"}).
		AddRow(1, "y", 1))

	_, err := FetchTicket(db, 1)
	assert.NoError(t, err)
}

func TestFetchTicketPrepareError(t *testing.T) {
	db := setUp()
	defer tearDown(db)

	sqlmock.ExpectPrepare().WillReturnError(fmt.Errorf("Some db error"))

	_, err := FetchTicket(db, 1)
	assert.Error(t, err)
}
