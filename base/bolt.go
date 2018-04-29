package base

import (
	"github.com/boltdb/bolt"
	"github.com/google/uuid"
	"github.com/nylar/triage/pkg/clock"
)

type Bolt struct {
	DB          *bolt.DB
	IDGenerator func() (uuid.UUID, error)
	Clock       clock.Clock
}
