package middleware

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func newLoggrus() {
	logger = logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
		DisableQuote:  true,
	})
}

func Infof(infoMsg string) {
	newLoggrus()
	logger.WithFields(logrus.Fields{}).Info(infoMsg)
}

func Errorf(err error) {
	newLoggrus()
	logger.WithFields(logrus.Fields{}).Errorf(fmt.Sprintf("Error : %s", err.Error()))
}
