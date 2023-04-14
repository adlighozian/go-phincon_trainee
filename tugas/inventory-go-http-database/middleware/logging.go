package middleware

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func Infof(infoMsg string) {
	logger = logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
		DisableQuote:  true,
	})
	logger.WithFields(logrus.Fields{}).Info(infoMsg)
}

func Errorf(err error) {
	logger = logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
		DisableQuote:  true,
	})
	logger.WithFields(logrus.Fields{}).Errorf(fmt.Sprintf("Error : %s", err.Error()))
}
