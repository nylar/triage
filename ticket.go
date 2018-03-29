package triage

import (
	"github.com/jmoiron/sqlx"
)

type Ticket struct {
	ID          int64   `json:"id"`
	Subject     string  `json:"subject"`
	Description *string `json:"description"`
	ProjectID   int64   `json:"project_id" db:"project_id"`
	StatusID    int64   `json:"status_id" db:"status_id"`
	TimeFields
}

func (t *Ticket) FindByID(db *sqlx.DB, id int64) error {
	query := `
SELECT
	id,
	subject,
	description,
	project_id,
        status_id,
	created_at,
	updated_at
FROM
	triage_ticket
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
        status_id,
	created_at,
	updated_at
FROM
	triage_ticket`

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
        status_id,
	created_at,
	updated_at
FROM
	triage_ticket
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
