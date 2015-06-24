package app

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestDb_Connect(t *testing.T) {
	u := "postgres"
	p := ""
	db := "triage_test"

	conn, err := Connect(u, p, db)
	defer conn.Close()

	assert.IsType(t, &sql.DB{}, conn)
	assert.NoError(t, err)
}
