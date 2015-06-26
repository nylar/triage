package api

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/nylar/triage/app"
	"github.com/stretchr/testify/assert"
)

var (
	conn   *sql.DB
	router *mux.Router
)

func init() {
	router = mux.NewRouter()
	router.StrictSlash(true)

	var err error
	conn, err = app.Connect("postgres", "", "triage")
	if err != nil {
		log.Fatalln(err.Error())
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
