package ticket

import (
	"context"
	"errors"

	"github.com/boltdb/bolt"
	"github.com/gogo/protobuf/proto"
	"github.com/nylar/triage/base"
	"github.com/nylar/triage/pkg/timeutil"
	"github.com/nylar/triage/ticket/ticketpb"
)

var (
	errBucketNotFound = errors.New("Triage bucket not found")
)

type Bolt struct {
	base.Bolt
}

func (bs *Bolt) List(ctx context.Context, req *ticketpb.ListRequest) (*ticketpb.ListResponse, error) {
	var tickets []*ticketpb.Ticket

	if err := bs.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(base.TicketBucket))
		if bucket == nil {
			return errBucketNotFound
		}

		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			ticket := &ticketpb.Ticket{}

			if err := proto.Unmarshal(v, ticket); err != nil {
				return err
			}

			tickets = append(tickets, ticket)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &ticketpb.ListResponse{
		Tickets: tickets,
	}, nil
}

func (bs *Bolt) Create(ctx context.Context, req *ticketpb.CreateRequest) (*ticketpb.CreateResponse, error) {
	id, err := bs.IDGenerator()
	if err != nil {
		return nil, err
	}

	now := bs.Clock.Now().UTC()

	ticket := &ticketpb.Ticket{
		Id:        id.String(),
		Subject:   req.Subject,
		CreatedAt: timeutil.TimeToTimestamp(now),
		UpdatedAt: timeutil.TimeToTimestamp(now),
	}

	if err := bs.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(base.TicketBucket))
		if bucket == nil {
			return errBucketNotFound
		}

		bytes, err := proto.Marshal(ticket)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(ticket.Id), bytes)
	}); err != nil {
		return nil, err
	}

	return &ticketpb.CreateResponse{
		Ticket: ticket,
	}, nil
}
