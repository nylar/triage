package app

import (
	"database/sql"
	"fmt"
)

// Connect takes a user, password and database to connect to a Postgres database
func Connect(user, password, database string) (*sql.DB, error) {
	connection := fmt.Sprintf("user=%s dbname=%s sslmode=disable",
		user, database)
	return sql.Open("postgres", connection)
}
