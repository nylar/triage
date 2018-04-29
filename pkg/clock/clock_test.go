package clock_test

import (
	"testing"

	"github.com/nylar/triage/pkg/clock"
)

func TestRealClock(t *testing.T) {
	var _ clock.Clock = clock.Real{}
}
