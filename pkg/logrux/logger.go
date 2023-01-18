package logrux

import (
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	logger := logrus.New()

	// Always use debug mode.
	logger.SetLevel(logrus.DebugLevel)

	logger.SetReportCaller(false)
	logger.SetFormatter(&logrus.TextFormatter{})
	return logger
}
