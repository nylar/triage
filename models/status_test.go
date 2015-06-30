package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatus_String(t *testing.T) {
	s := new(Status)
	s.Name = "open"

	assert.Equal(t, s.String(), "open")
}

func TestStatus_DefaultStatus(t *testing.T) {
	s := DefaultStatus()
	assert.Equal(t, s.Name, "open")
	assert.Equal(t, s.StatusID, 1)
}
