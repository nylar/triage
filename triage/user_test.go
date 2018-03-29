package triage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserFindByID(t *testing.T) {
	teardown := setUp(t)
	defer teardown()

	loadFixtures(t)

	user := &User{}

	err := user.FindByID(db, 1)

	assert.NoError(t, err)
}
