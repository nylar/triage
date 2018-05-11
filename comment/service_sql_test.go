package comment_test

import (
	"testing"

	"github.com/nylar/triage/ticket"
)

func TestSQLService(t *testing.T) {
	var _ ticket.Service = &ticket.SQL{}
}
