package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModels_Ticket_NewTicket(t *testing.T) {
	tkt := NewTicket("")
	assert.IsType(t, &Ticket{}, tkt)
}

func TestModels_Ticket_String(t *testing.T) {
	msg := "I need help!"
	tkt := NewTicket(msg)

	assert.Equal(t, tkt.String(), msg)
}
