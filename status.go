package triage

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Status struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Colour    string    `json:"colour"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Statuses struct {
	Statuses []*Status `json:"statuses"`
}

func (s *Statuses) FindAll(db *sqlx.DB) error {
	query := `
SELECT
	id,
	name,
	colour,
	created_at,
	updated_at
FROM
	status`

	rows, err := db.Queryx(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		status := &Status{}
		if err := rows.StructScan(status); err != nil {
			return err
		}

		s.Statuses = append(s.Statuses, status)
	}

	return nil
}
