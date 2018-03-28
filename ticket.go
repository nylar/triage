package triage

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Ticket struct {
	ID          int64     `json:"id"`
	Subject     string    `json:"subject"`
	Description *string   `json:"description"`
	ProjectID   int64     `json:"project_id" db:"project_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func (t *Ticket) FindByID(db *sqlx.DB, id int64) error {
	query := `
SELECT
	id,
	subject,
	description,
	project_id,
	created_at,
	updated_at
FROM
	ticket
WHERE
	id = ?`

	return db.Get(t, query, id)
}

type Tickets struct {
	Tickets []*Ticket `json:"tickets"`
}

func (t *Tickets) FindAll(db *sqlx.DB) error {
	query := `
SELECT
	id,
	subject,
	description,
	project_id,
	created_at,
	updated_at
FROM
	ticket`

	rows, err := db.Queryx(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		ticket := &Ticket{}
		if err := rows.StructScan(ticket); err != nil {
			return err
		}

		t.Tickets = append(t.Tickets, ticket)
	}

	return nil
}

func (t *Tickets) FindByProjectID(db *sqlx.DB, projectID int64) error {
	query := `
SELECT
	id,
	subject,
	description,
	project_id,
	created_at,
	updated_at
FROM
	ticket
WHERE
	project_id = ?`

	rows, err := db.Queryx(query, projectID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		ticket := &Ticket{}
		if err := rows.StructScan(ticket); err != nil {
			return err
		}

		t.Tickets = append(t.Tickets, ticket)
	}

	return nil
}
