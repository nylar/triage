package comment

import (
	"context"

	"github.com/elgris/sqrl"
	"github.com/go-sql-driver/mysql"
	"github.com/nylar/triage/base"
	"github.com/nylar/triage/comment/commentpb"
	"github.com/nylar/triage/pkg/timeutil"
)

const tableName = "triage_comment"

type SQL struct {
	base.SQL
}

func (cs *SQL) List(ctx context.Context, req *commentpb.ListRequest) (*commentpb.ListResponse, error) {
	where := sqrl.Eq{}
	if req.TicketId != "" {
		where["ticket_id"] = req.TicketId
	}

	rows, err := sqrl.
		Select("id", "ticket_id", "content", "created_at", "updated_at").
		From(tableName).
		Where(where).
		PlaceholderFormat(cs.Placeholder).
		RunWith(cs.DB).
		QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*commentpb.Comment

	var createdAt, updatedAt mysql.NullTime

	for rows.Next() {
		comment := &commentpb.Comment{}

		if err := rows.Scan(
			&comment.Id,
			&comment.TicketId,
			&comment.Content,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, err
		}

		comment.CreatedAt = timeutil.TimeToTimestamp(createdAt.Time)
		comment.UpdatedAt = timeutil.TimeToTimestamp(updatedAt.Time)

		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &commentpb.ListResponse{
		Comments: comments,
	}, nil
}

func (cs *SQL) Create(ctx context.Context, req *commentpb.CreateRequest) (*commentpb.CreateResponse, error) {
	id, err := cs.IDGenerator()
	if err != nil {
		return nil, err
	}

	now := timeutil.TimeToTimestamp(cs.Clock.Now().UTC())

	comment := &commentpb.Comment{
		Id:        id.String(),
		TicketId:  req.TicketId,
		Content:   req.Content,
		CreatedAt: now,
		UpdatedAt: now,
	}

	_, err = sqrl.
		Insert(tableName).
		Columns("id", "ticket_id", "content", "created_at", "updated_at").
		Values(
			comment.Id,
			comment.TicketId,
			comment.Content,
			timeutil.TimestampToTime(comment.CreatedAt),
			timeutil.TimestampToTime(comment.UpdatedAt),
		).
		PlaceholderFormat(cs.Placeholder).
		RunWith(cs.DB).
		ExecContext(ctx)
	if err != nil {
		return nil, err
	}

	return &commentpb.CreateResponse{
		Comment: comment,
	}, nil
}
