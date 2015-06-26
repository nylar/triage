package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAsInt(t *testing.T) {
	num, err := ParseToInt("1")
	assert.Equal(t, num, 1)
	assert.NoError(t, err)

	num, err = ParseToInt("987654321")
	assert.Equal(t, num, 987654321)
	assert.NoError(t, err)

	num, err = ParseToInt("xx")
	assert.Equal(t, num, 0)
	assert.Error(t, err)
}

func TestParseAsInt64(t *testing.T) {
	expectedType := new(int64)

	num, err := ParseToInt64("1")
	assert.Equal(t, num, 1)
	assert.NoError(t, err)
	assert.IsType(t, num, *expectedType)
}
