package ticket_test

import (
	"time"

	"github.com/nylar/triage/ticket/ticketpb"
)

// Fixtures
var ticketFixtures = []*ticketpb.Ticket{
	&ticketpb.Ticket{
		Id:      "1",
		Subject: "My first ticket",
	},
}

// Mocks
type mockClock struct {
	t time.Time
}

func (mc mockClock) Now() time.Time {
	return mc.t
}
