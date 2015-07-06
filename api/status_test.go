package api

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestStatusIndex(t *testing.T) {
	db := setUp()
	defer tearDown(db)

	sqlmock.ExpectQuery("SELECT (.+) FROM status").
		WillReturnRows(sqlmock.
		NewRows([]string{"status_id", "name"}).
		AddRow(1, "open"))

	route := server.URL + "/api/statuses"
	resp, err := http.Get(route)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestStatusIndexError(t *testing.T) {
	db := setUp()
	defer tearDown(db)

	sqlmock.ExpectQuery("SELECT (.+) FROM status").
		WillReturnError(fmt.Errorf("Query failed"))

	route := server.URL + "/api/statuses"
	resp, err := http.Get(route)

	assert.NoError(t, err)
	assert.Equal(t, 500, resp.StatusCode)
}

func TestFetchStatuses(t *testing.T) {
	db := setUp()
	defer tearDown(db)

	sqlmock.ExpectQuery("SELECT (.+) FROM status").
		WillReturnRows(sqlmock.
		NewRows([]string{"status_id", "name"}).
		AddRow(1, "open"))

	_, err := FetchStatuses(db)
	assert.NoError(t, err)
}

func TestFetchStatusesRowsError(t *testing.T) {
	db := setUp()
	defer tearDown(db)

	sqlmock.ExpectQuery("SELECT (.+) FROM status").
		WillReturnError(fmt.Errorf("Query failed"))

	_, err := FetchStatuses(db)
	assert.Error(t, err)
}

func TestFetchStatusesScanError(t *testing.T) {
	db := setUp()
	defer tearDown(db)

	sqlmock.ExpectQuery("SELECT (.+) FROM status").
		WillReturnRows(sqlmock.
		NewRows([]string{"status_id", "name"}).
		AddRow("x", "y"))

	_, err := FetchStatuses(db)
	assert.Error(t, err)
}
