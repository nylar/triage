package main

type Config struct {
	Server Server `toml:"server"`
}

type Server struct {
	Address string `toml:"address"`
}
