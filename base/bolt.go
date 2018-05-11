package base

import (
	"github.com/boltdb/bolt"
	"github.com/google/uuid"
	"github.com/nylar/triage/pkg/clock"
)

// Buckets
const (
	TicketBucket         = "ticket"
	CommentBucket        = "comment"
	TicketCommentsBucket = "ticket_comments"
)

// BoltRecordSeparator is the 'information separator one' character, ASCII 31
const BoltRecordSeparator = '\u001F'

var buckets = []string{
	TicketBucket,
	CommentBucket,
	TicketCommentsBucket,
}

type Bolt struct {
	DB          *bolt.DB
	IDGenerator func() (uuid.UUID, error)
	Clock       clock.Clock
}

// Bootstrap ensures the required buckets are created
func (bs *Bolt) Bootstrap() error {
	return bs.DB.Update(func(tx *bolt.Tx) error {
		for _, bucket := range buckets {
			_, err := tx.CreateBucketIfNotExists([]byte(bucket))
			if err != nil {
				return err
			}
		}
		return nil
	})
}
