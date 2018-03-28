package triage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTicketFindByID(t *testing.T) {
	teardown := setUp(t)
	defer teardown()

	loadFixtures(t)

	ticket := &Ticket{}

	err := ticket.FindByID(db, 1)

	assert.NoError(t, err)
}

func TestTicketFindAll(t *testing.T) {
	teardown := setUp(t)
	defer teardown()

	loadFixtures(t)

	tickets := &Tickets{}

	err := tickets.FindAll(db)

	assert.NoError(t, err)

	assert.Equal(t, 2, len(tickets.Tickets), "Expected two tickets")
}

func TestTicketFindByProjectID(t *testing.T) {
	teardown := setUp(t)
	defer teardown()

	loadFixtures(t)

	tickets := &Tickets{}

	err := tickets.FindByProjectID(db, 1)

	assert.NoError(t, err)

	assert.Equal(t, 2, len(tickets.Tickets), "Expected two tickets")

	tickets2 := &Tickets{}

	err = tickets2.FindByProjectID(db, 9999)

	assert.NoError(t, err)

	assert.Equal(t, 0, len(tickets2.Tickets), "Expected zero tickets")
}
