package ticket_test

import "github.com/nylar/triage/ticket/ticketpb"

// Fixtures
var ticketFixtures = []*ticketpb.Ticket{
	&ticketpb.Ticket{
		Id:      "1",
		Subject: "My first ticket",
	},
}
