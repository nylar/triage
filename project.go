package triage

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Project struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// FindByID returns an individual project. A missing project can be determined
// by checking the error for an sql.ErrNoRows.
func (p *Project) FindByID(db *sqlx.DB, id int64) error {
	query := `
SELECT
	id,
	name,
	created_at,
	updated_at
FROM
	project
WHERE
	id = ?`

	return db.Get(p, query, id)
}

func (p *Project) Create(db *sqlx.DB) error {
	res, err := db.Exec(`INSERT INTO project (name) VALUES (?)`, p.Name)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	p.ID = id
	return nil
}

type Projects struct {
	Projects []*Project `json:"projects"`
}

func (p *Projects) FindAll(db *sqlx.DB) error {
	query := `
SELECT
	id,
	name,
	created_at,
	updated_at
FROM
	project`

	rows, err := db.Queryx(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		project := &Project{}
		if err := rows.StructScan(project); err != nil {
			return err
		}

		p.Projects = append(p.Projects, project)
	}

	return nil
}
