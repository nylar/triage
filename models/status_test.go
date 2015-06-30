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
