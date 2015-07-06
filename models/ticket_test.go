package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModels_Ticket_NewTicket(t *testing.T) {
	tkt := NewTicket("")
	assert.IsType(t, &Ticket{}, tkt)
	assert.Equal(t, tkt.Status.Name, "open")
}

func TestModels_Ticket_String(t *testing.T) {
	subject := "I need help!"
	tkt := NewTicket(subject)

	assert.Equal(t, tkt.String(), subject)
}
