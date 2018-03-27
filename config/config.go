package config

import (
	"fmt"
	"net/url"
)

// HTTP server settings
type HTTP struct {
	Port     int    `toml:"port"`
	Hostname string `toml:"hostname"`
}

// SQL database settings
type SQL struct {
	Hostname string            `toml:"hostname"`
	Database string            `toml:"database"`
	Username string            `toml:"username"`
	Password string            `toml:"password"`
	Port     int               `toml:"port"`
	Params   map[string]string `toml:"params"`
}

// DataSourceName builds the connection string
func (s SQL) DataSourceName() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		s.Username,
		s.Password,
		s.Hostname,
		s.Port,
		s.Database,
		s.encodeParams(),
	)
}

func (s SQL) encodeParams() string {
	values := url.Values{}

	for key, value := range s.Params {
		values.Set(key, value)
	}

	return values.Encode()
}

// Config holds settings for the application
type Config struct {
	HTTP HTTP `toml:"http"`
	SQL  SQL  `toml:"sql"`
}
