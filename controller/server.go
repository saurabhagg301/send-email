package controller

import (
	"net/http"
	"time"

	"github.com/saurabh-arch/send-email/logger"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logger.GetLogger()
}

// StartWebServer to start webserver
func StartWebServer() {
	// get router
	r := GetRouter()

	// webserver configuration
	srvr := http.Server{
		Addr:         "localhost:8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 60 * time.Second,
		Handler:      r,
	}

	// start webserver
	log.Fatal(srvr.ListenAndServe())
}
