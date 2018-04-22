package ticket

import (
	"context"
	"database/sql"

	"github.com/elgris/sqrl"
	"github.com/google/uuid"
	"github.com/nylar/triage/ticket/ticketpb"
)

const tableName = "triage_ticket"

// SQL facilitates the management of tickets, backed by an SQL database.
type SQL struct {
	DB          *sql.DB
	Placeholder sqrl.PlaceholderFormat
	IDGenerator func() (uuid.UUID, error)
}

// List returns a list of zero or more tickets that match the criteria in the
// request
func (ts *SQL) List(ctx context.Context, req *ticketpb.ListRequest) (*ticketpb.ListResponse, error) {
	rows, err := sqrl.
		Select("id", "subject").
		From(tableName).
		PlaceholderFormat(ts.Placeholder).
		RunWith(ts.DB).
		Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []*ticketpb.Ticket

	for rows.Next() {
		ticket := &ticketpb.Ticket{}

		if err := rows.Scan(&ticket.Id, &ticket.Subject); err != nil {
			return nil, err
		}

		tickets = append(tickets, ticket)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &ticketpb.ListResponse{
		Tickets: tickets,
	}, nil
}

func (ts *SQL) Create(ctx context.Context, req *ticketpb.CreateRequest) (*ticketpb.CreateResponse, error) {
	id, err := ts.IDGenerator()
	if err != nil {
		return nil, err
	}

	ticket := &ticketpb.Ticket{
		Id:      id.String(),
		Subject: req.Subject,
	}

	_, err = sqrl.
		Insert(tableName).
		Columns("id", "subject").
		Values(ticket.Id, ticket.Subject).
		PlaceholderFormat(ts.Placeholder).
		RunWith(ts.DB).
		Exec()
	if err != nil {
		return nil, err
	}

	return &ticketpb.CreateResponse{
		Ticket: ticket,
	}, nil
}
