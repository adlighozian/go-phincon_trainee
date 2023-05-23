package logging

import (
	"fmt"
	"net/http"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func newLoggrus() {
	log = logrus.New()
	// Will log anything that is info or above (warn, error, fatal, panic)
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
		DisableQuote:  true,
	})
}

func Infof(infoMsg string, r *http.Request) {
	fields := logrus.Fields{}
	if r.Method != "" {
		fields["Host"] = r.Host
		fields["Method"] = r.Method
		fields["URL"] = r.RequestURI
	}

	newLoggrus()
	log.WithFields(fields).Info(infoMsg)
}

func Errorf(err error, r *http.Request) {
	fields := logrus.Fields{}
	if r.Method != "" {
		fields["Host"] = r.Host
		fields["Method"] = r.Method
		fields["URL"] = r.RequestURI
	}
	if r.Body != nil {
		fields["Params"] = r.Body
	}

	newLoggrus()
	log.WithFields(fields).Errorf(fmt.Sprintf("Error : %s", err.Error()))
}
