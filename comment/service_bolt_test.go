package comment_test

import (
	"testing"

	"github.com/nylar/triage/ticket"
)

func TestBoltService(t *testing.T) {
	var _ ticket.Service = &ticket.Bolt{}
}
