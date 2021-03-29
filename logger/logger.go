package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// GetLogger function to get a logger instance
func GetLogger() *logrus.Logger {
	// Create a new instance of the logger
	var log = logrus.New()

	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	log.Out = os.Stdout

	// You could set this to any `io.Writer` such as a file
	// file, err := os.OpenFile("send-email.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err == nil {
	//  log.Out = file
	// } else {
	//  log.Info("Failed to log to file, using default stderr")
	// }

	// // To log to /var/log/syslog file (in Linux) or /var/log/system.log file (in MacOS)
	// logWriter, err := syslog.New(syslog.LOG_NOTICE, "(send-email)")
	// if err == nil {
	// 	log.Out = logWriter
	// }

	return log
}
