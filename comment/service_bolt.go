package comment

import (
	"bytes"
	"context"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/gogo/protobuf/proto"
	"github.com/nylar/triage/base"
	"github.com/nylar/triage/comment/commentpb"
	"github.com/nylar/triage/pkg/timeutil"
)

type Bolt struct {
	base.Bolt
}

func (bs *Bolt) List(ctx context.Context, req *commentpb.ListRequest) (*commentpb.ListResponse, error) {
	var comments []*commentpb.Comment

	if err := bs.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(base.TicketCommentsBucket))
		if bucket == nil {
			return fmt.Errorf("%s bucket not found", base.TicketCommentsBucket)
		}

		value := bucket.Get([]byte(req.TicketId))
		if value == nil {
			return nil
		}

		ids := bytes.Split(value, []byte{base.BoltRecordSeparator})

		bucket = tx.Bucket([]byte(base.CommentBucket))
		if bucket == nil {
			return fmt.Errorf("%s bucket not found", base.CommentBucket)
		}

		for _, id := range ids {
			comment := &commentpb.Comment{}

			bytes := bucket.Get(id)
			if bytes == nil {
				return fmt.Errorf("Could not find comment '%s'", string(id))
			}

			if err := proto.Unmarshal(bytes, comment); err != nil {
				return err
			}

			comments = append(comments, comment)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &commentpb.ListResponse{
		Comments: comments,
	}, nil
}

func (bs *Bolt) Create(ctx context.Context, req *commentpb.CreateRequest) (*commentpb.CreateResponse, error) {
	id, err := bs.IDGenerator()
	if err != nil {
		return nil, err
	}

	now := bs.Clock.Now().UTC()

	comment := &commentpb.Comment{
		Id:        id.String(),
		TicketId:  req.TicketId,
		Content:   req.Content,
		CreatedAt: timeutil.TimeToTimestamp(now),
		UpdatedAt: timeutil.TimeToTimestamp(now),
	}

	if err := bs.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(base.CommentBucket))
		if bucket == nil {
			return fmt.Errorf("%s bucket not found", base.CommentBucket)
		}

		bytes, err := proto.Marshal(comment)
		if err != nil {
			return err
		}

		if err := bucket.Put([]byte(comment.Id), bytes); err != nil {
			return err
		}

		bucket = tx.Bucket([]byte(base.TicketCommentsBucket))
		if bucket == nil {
			return fmt.Errorf("%s bucket not found", base.TicketCommentsBucket)
		}

		ids := bucket.Get([]byte(comment.TicketId))
		if ids == nil {
			ids = []byte(comment.Id)
		} else {
			ids = append(ids, base.BoltRecordSeparator)
			ids = append(ids, comment.Id...)
		}

		return bucket.Put([]byte(comment.TicketId), ids)
	}); err != nil {
		return nil, err
	}

	return &commentpb.CreateResponse{
		Comment: comment,
	}, nil
}
