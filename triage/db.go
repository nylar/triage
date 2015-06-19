package triage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect(user, password, database string) (*sql.DB, error) {
	connection := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		user, password, database)
	return sql.Open("postgres", connection)
}
