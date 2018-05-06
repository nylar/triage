package config_test

import (
	"testing"

	"github.com/elgris/sqrl"
	"github.com/nylar/triage/config"
	"github.com/stretchr/testify/assert"
)

func TestSqlDsn(t *testing.T) {
	tests := []struct {
		name        string
		sql         config.SQL
		expected    string
		shouldError bool
	}{
		{
			name: "mysql protocol",
			sql: config.SQL{
				Vendor:   "mysql",
				Host:     "localhost",
				Port:     3306,
				Username: "root",
				Password: "secret",
				Database: "db",
				Params: map[string]string{
					"foo": "bar",
				},
			},
			expected:    "root:secret@tcp(localhost:3306)/db?foo=bar",
			shouldError: false,
		},
		{
			name: "postgres protocol",
			sql: config.SQL{
				Vendor:   "postgres",
				Host:     "localhost",
				Port:     3306,
				Username: "root",
				Password: "secret",
				Database: "db",
				Params: map[string]string{
					"foo": "bar",
				},
			},
			expected:    "postgres://root:secret@localhost:3306/db?foo=bar",
			shouldError: false,
		},
		{
			name: "unsupported database",
			sql: config.SQL{
				Vendor: "foobar",
			},
			expected:    "",
			shouldError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dsn, err := test.sql.DSN()

			assert.Equal(t, test.expected, dsn)

			if test.shouldError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSqlPlaceholder(t *testing.T) {
	tests := []struct {
		name     string
		sql      config.SQL
		expected sqrl.PlaceholderFormat
	}{
		{
			name: "mysql protocol",
			sql: config.SQL{
				Vendor: "mysql",
			},
			expected: sqrl.Question,
		},
		{
			name: "postgres protocol",
			sql: config.SQL{
				Vendor: "postgres",
			},
			expected: sqrl.Dollar,
		},
		{
			name: "unsupported database",
			sql: config.SQL{
				Vendor: "foobar",
			},
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			placeholder := test.sql.Placeholder()

			assert.Equal(t, test.expected, placeholder)
		})
	}
}

func TestServerAddress(t *testing.T) {
	server := config.Server{
		Port: 3000,
	}

	assert.Equal(t, ":3000", server.Address())
}
