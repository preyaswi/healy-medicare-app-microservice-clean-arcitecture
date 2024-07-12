package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.DebugLevel)
}
func Logger() *logrus.Logger {
	return log
}
