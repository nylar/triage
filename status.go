package triage

import (
	"github.com/jmoiron/sqlx"
)

type Status struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Colour string `json:"colour"`
	TimeFields
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
	triage_status`

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
