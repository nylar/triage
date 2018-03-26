package main

import (
	"fmt"

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
}

func version() string {
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}
