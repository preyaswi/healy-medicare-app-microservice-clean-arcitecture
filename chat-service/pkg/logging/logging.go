package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Init() {
	// Set the output of the log to stdout
	log.SetOutput(os.Stdout)

	// Set the log level (you can change this to logrus.WarnLevel or logrus.ErrorLevel as needed)
	log.SetLevel(logrus.DebugLevel)

	// Set the formatter to JSON for structured logging
	log.SetFormatter(&logrus.JSONFormatter{})
}

func Logger() *logrus.Logger {
	return log
}
