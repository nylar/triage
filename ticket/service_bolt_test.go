package ticket_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/boltdb/bolt"
	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	"github.com/nylar/triage/base"
	"github.com/nylar/triage/ticket"
	"github.com/nylar/triage/ticket/ticketpb"
	"github.com/stretchr/testify/assert"
)

var mockUUIDGenerator = func() (uuid.UUID, error) {
	return uuid.Parse("9958d76c-6e3d-4eef-911b-175b39e8f696")
}

func tempfile() string {
	f, _ := ioutil.TempFile("", "")
	f.Close()
	os.Remove(f.Name())
	return f.Name()
}

func setUpBolt(t *testing.T) *bolt.DB {
	t.Helper()

	db, err := bolt.Open(tempfile(), 0666, nil)
	if err != nil {
		t.Fatalf("Couldn't open bolt DB: %v", err)
	}

	return db
}

func teardownBolt(t *testing.T, db *bolt.DB) {
	t.Helper()

	defer os.Remove(db.Path())

	if err := db.Close(); err != nil {
		t.Fatalf("Couldn't close bolt DB: %v", err)
	}
}

func loadBoltFixtures(t *testing.T, db *bolt.DB) {
	t.Helper()

	if err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("ticket"))
		if bucket == nil {
			return fmt.Errorf("Couldn't find bucket")
		}

		for _, fixture := range ticketFixtures {
			bytes, err := proto.Marshal(fixture)
			if err != nil {
				return err
			}

			if err := bucket.Put([]byte(fixture.Id), bytes); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		t.Fatalf("Couldn't load fixtures: %v", err)
	}
}

func TestBoltService(t *testing.T) {
	var _ ticket.Service = &ticket.Bolt{}
}

func TestBoltCreate(t *testing.T) {
	db := setUpBolt(t)
	defer teardownBolt(t, db)

	service := &ticket.Bolt{
		Bolt: base.Bolt{
			DB:          db,
			IDGenerator: mockUUIDGenerator,
			Clock:       mockClock{time.Date(2018, time.July, 25, 7, 30, 45, 0, time.UTC)},
		},
	}

	if err := service.Bootstrap(); err != nil {
		t.Fatalf("Couldn't bootstrap DB: %v", err)
	}

	subject := "New ticket"

	res, err := service.Create(context.Background(), &ticketpb.CreateRequest{
		Subject: subject,
	})

	assert.NoError(t, err)
	assert.Equal(t, subject, res.Ticket.Subject)
}

func TestBoltList(t *testing.T) {
	db := setUpBolt(t)
	defer teardownBolt(t, db)

	service := &ticket.Bolt{
		Bolt: base.Bolt{
			DB: db,
		},
	}

	if err := service.Bootstrap(); err != nil {
		t.Fatalf("Couldn't bootstrap DB: %v", err)
	}

	loadBoltFixtures(t, db)

	res, err := service.List(context.Background(), &ticketpb.ListRequest{})

	assert.NoError(t, err)
	assert.Equal(t, ticketFixtures, res.Tickets)
}
