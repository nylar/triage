package ticket

import (
	"context"

	"github.com/nylar/triage/ticket/ticketpb"
)

// Service provides methods for retrieving ticket(s) from a datastore
type Service interface {
	List(context.Context, *ticketpb.ListRequest) (*ticketpb.ListResponse, error)
	Create(context.Context, *ticketpb.CreateRequest) (*ticketpb.CreateResponse, error)
}
