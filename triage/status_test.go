package triage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusFindAll(t *testing.T) {
	teardown := setUp(t)
	defer teardown()

	loadFixtures(t)

	statuses := &Statuses{}

	err := statuses.FindAll(db)

	assert.NoError(t, err)

	assert.Equal(t, 3, len(statuses.Statuses), "Expected three statuses")
}
