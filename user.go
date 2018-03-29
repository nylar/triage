package triage

import "github.com/jmoiron/sqlx"

type User struct {
	ID             int64  `json:"id"`
	Username       string `json:"username"`
	HashedPassword []byte `json:"-" db:"hashed_password"`
	TimeFields
}

func (u *User) FindByID(db *sqlx.DB, id int64) error {
	query := `
SELECT
	id,
	username,
    hashed_password,
	created_at,
	updated_at
FROM
	triage_user
WHERE
	id = ?`

	return db.Get(u, query, id)
}
