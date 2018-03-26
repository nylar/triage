package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/nylar/triage/config"
	toml "github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
)

const (
	major = 0
	minor = 1
	patch = 0
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableSorting: true,
	})
}

func main() {
	logrus.WithField("version", version()).Infoln("Triage")

	configFile, err := os.Open(os.Getenv("TRIAGE_CONFIG_PATH"))
	if err != nil {
		logrus.WithError(err).Fatalln("Couldn't open config file")
	}

	conf := &config.Config{}
	if err := toml.NewDecoder(configFile).Decode(conf); err != nil {
		logrus.WithError(err).Fatalln("Couldn't parse config file")
	}

	server := &http.Server{
		Addr: fmt.Sprintf("%s:%d", conf.HTTP.Hostname, conf.HTTP.Port),
	}
	idleConnections := make(chan struct{})
	go shutdown(server, idleConnections)

	logrus.WithField("port", conf.HTTP.Port).Infoln("HTTP server starting")

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logrus.WithError(err).Fatalln("HTTP server failed to start")
	}
	<-idleConnections
}

func shutdown(server *http.Server, idleConnections chan struct{}) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.WithError(err).Fatalln("HTTP server shutdown failure")
	}

	logrus.Infoln("HTTP server shutdown successful")
	close(idleConnections)
}

func version() string {
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}
