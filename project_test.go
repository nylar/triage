package triage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func projectFixture(t *testing.T) {
	project := &Project{
		ID:   1,
		Name: "Test Project",
	}

	_, err := db.Exec(
		`INSERT INTO project (id, name) VALUES (?, ?)`,
		project.ID,
		project.Name,
	)

	if err != nil {
		t.Fatalf("Couldn't insert project fixture: %v", err)
	}
}

func TestProjectFindByID(t *testing.T) {
	teardown := setUp(t)
	defer teardown()

	// Load fixtures
	projectFixture(t)

	project := &Project{}

	err := project.FindByID(db, 1)

	assert.NoError(t, err)
}
