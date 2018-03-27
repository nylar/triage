package triage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectFindByID(t *testing.T) {
	teardown := setUp(t)
	defer teardown()

	loadFixtures(t)

	project := &Project{}

	err := project.FindByID(db, 1)

	assert.NoError(t, err)
}

func TestProjectsFindAll(t *testing.T) {
	teardown := setUp(t)
	defer teardown()

	loadFixtures(t)

	projects := &Projects{}

	err := projects.FindAll(db)

	assert.NoError(t, err)

	assert.Equal(t, 2, len(projects.Projects), "Expected two projects")
}
