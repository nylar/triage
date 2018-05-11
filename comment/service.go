package comment

import (
	"context"

	"github.com/nylar/triage/comment/commentpb"
)

type Service interface {
	List(context.Context, *commentpb.ListRequest) (*commentpb.ListResponse, error)
	Create(context.Context, *commentpb.CreateRequest) (*commentpb.CreateResponse, error)
}
