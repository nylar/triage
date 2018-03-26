package triage

import "github.com/jmoiron/sqlx"

type Project struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// FindByID returns an individual project. A missing project can be determined
// by checking the error for an sql.ErrNoRows.
func (p *Project) FindByID(db *sqlx.DB, id int64) error {
	return db.Get(p, `SELECT id, name FROM project WHERE id = ?`, id)
}
