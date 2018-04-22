package config

import (
	"fmt"
	"net/url"

	"github.com/elgris/sqrl"
)

// Vendors
const (
	mysql    = "mysql"    // MySQL, MariaDB
	postgres = "postgres" // Postgres, Cockroach
)

// Config groups each configuration into a top-level object
type Config struct {
	Server Server `toml:"server"`
	SQL    *SQL   `toml:"sql"`
	Bolt   *Bolt  `toml:"bolt"`
}

// Server provides configuration for the gRPC server
type Server struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

// Address builds the address that the server will listen on
func (s Server) Address() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

// SQL provides configuration for connecting to an SQL database
type SQL struct {
	Vendor   string            `toml:"vendor"`
	Host     string            `toml:"host"`
	Port     int               `toml:"port"`
	Username string            `toml:"username"`
	Password string            `toml:"password"`
	Database string            `toml:"database"`
	Params   map[string]string `toml:"params"`
}

// DSN builds the connection string for each supported database
func (s SQL) DSN() (string, error) {
	switch s.Vendor {
	case mysql:
		return s.mysqlDSN(), nil
	case postgres:
		return s.postgresDSN(), nil
	default:
		return "", fmt.Errorf("Unable to handle '%s'", s.Vendor)
	}
}

func (s SQL) mysqlDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		s.Username,
		s.Password,
		s.Host,
		s.Port,
		s.Database,
		s.encodedParams(),
	)
}

func (s SQL) postgresDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?%s",
		s.Username,
		s.Password,
		s.Host,
		s.Port,
		s.Database,
		s.encodedParams(),
	)
}

func (s SQL) encodedParams() string {
	params := url.Values{}

	for key, value := range s.Params {
		params.Set(key, value)
	}

	return params.Encode()
}

// Placeholder determines which placeholder delimiter will be used when
// constructing queries
func (s SQL) Placeholder() sqrl.PlaceholderFormat {
	switch s.Vendor {
	case mysql:
		return sqrl.Question
	case postgres:
		return sqrl.Dollar
	default:
		return nil
	}
}

type Bolt struct {
	Path string `toml:"path"`
}
