package comment_test

import (
	"time"

	"github.com/nylar/triage/base"
	"github.com/nylar/triage/comment/commentpb"
)

// Fixtures
var commentFixtures = map[string]*commentpb.Comment{
	base.CommentBucket: &commentpb.Comment{
		Id:       "1",
		TicketId: "1",
		Content:  "My first comment",
	},
}

// Mocks
type mockClock struct {
	t time.Time
}

func (mc mockClock) Now() time.Time {
	return mc.t
}
