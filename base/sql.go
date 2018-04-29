package base

import (
	"database/sql"

	"github.com/elgris/sqrl"
	"github.com/google/uuid"
	"github.com/nylar/triage/pkg/clock"
)

type SQL struct {
	DB          *sql.DB
	Placeholder sqrl.PlaceholderFormat
	IDGenerator func() (uuid.UUID, error)
	Clock       clock.Clock
}
