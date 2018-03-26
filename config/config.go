package config

// HTTP server settings
type HTTP struct {
	Port     int    `toml:"port"`
	Hostname string `toml:"hostname"`
}

type Config struct {
	HTTP HTTP `toml:"http"`
}
