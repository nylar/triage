package api

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var server *httptest.Server

func setUp() *sql.DB {
	mockdb, err := sqlmock.New()
	if err != nil {
		log.Fatalln(err.Error())
	}

	routes := httprouter.New()
	Routes(routes, mockdb)
	server = httptest.NewServer(routes)

	return mockdb
}

func tearDown(db *sql.DB) {
	server.Close()

	if err := db.Close(); err != nil {
		log.Fatalf("Error '%s' was not expected while closing the database", err)
	}
}

func TestErrorResponse(t *testing.T) {
	w := httptest.NewRecorder()
	err := "I am a teapot"

	errorResponse(w, err, http.StatusTeapot)

	assert.Equal(
		t,
		"{\"status_code\":418,\"message\":\"I am a teapot\"}\n",
		w.Body.String(),
	)
}
