package ticket

import (
	"context"

	"github.com/elgris/sqrl"
	"github.com/go-sql-driver/mysql"
	"github.com/nylar/triage/base"
	"github.com/nylar/triage/pkg/timeutil"
	"github.com/nylar/triage/ticket/ticketpb"
)

const tableName = "triage_ticket"

// SQL facilitates the management of tickets, backed by an SQL database.
type SQL struct {
	base.SQL
}

// List returns a list of zero or more tickets that match the criteria in the
// request
func (ts *SQL) List(ctx context.Context, req *ticketpb.ListRequest) (*ticketpb.ListResponse, error) {
	rows, err := sqrl.
		Select("id", "subject", "created_at", "updated_at").
		From(tableName).
		PlaceholderFormat(ts.Placeholder).
		RunWith(ts.DB).
		Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []*ticketpb.Ticket

	var createdAt, updatedAt mysql.NullTime

	for rows.Next() {
		ticket := &ticketpb.Ticket{}

		if err := rows.Scan(
			&ticket.Id,
			&ticket.Subject,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, err
		}

		ticket.CreatedAt = timeutil.TimeToTimestamp(createdAt.Time)
		ticket.UpdatedAt = timeutil.TimeToTimestamp(updatedAt.Time)

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

	now := ts.Clock.Now().UTC()

	ticket := &ticketpb.Ticket{
		Id:        id.String(),
		Subject:   req.Subject,
		CreatedAt: timeutil.TimeToTimestamp(now),
		UpdatedAt: timeutil.TimeToTimestamp(now),
	}

	_, err = sqrl.
		Insert(tableName).
		Columns("id", "subject", "created_at", "updated_at").
		Values(
			ticket.Id,
			ticket.Subject,
			timeutil.TimestampToTime(ticket.CreatedAt),
			timeutil.TimestampToTime(ticket.UpdatedAt)).
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
