package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
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

	r, err := http.NewRequest("GET", "/api/statuses/", nil)
	if err != nil {
		t.Error(err.Error())
	}

	w := httptest.NewRecorder()
	h := StatusIndex(db)
	h.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code)

	Routes(router, db)
	router.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code)
}

func TestStatusIndexError(t *testing.T) {
	db := setUp()
	defer tearDown(db)

	sqlmock.ExpectQuery("SELECT (.+) FROM status").
		WillReturnError(fmt.Errorf("Query failed"))

	r, err := http.NewRequest("GET", "/api/statuses/", nil)
	if err != nil {
		t.Error(err.Error())
	}

	w := httptest.NewRecorder()
	h := StatusIndex(db)
	h.ServeHTTP(w, r)

	assert.Equal(t, 500, w.Code)
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
