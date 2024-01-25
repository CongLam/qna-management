package utils

import (
	"time"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Logger
)

func init() {
	log = logrus.New()
	log.SetLevel(logrus.InfoLevel)

	// Get the current date to use in the log file name
	currentDate := time.Now().Format("2006-01-02")

	// Create a new instance of the lfshook with daily log files
	logFilePath := "logs/" + currentDate + ".log"
	lfHook := lfshook.NewHook(
		lfshook.PathMap{
			logrus.InfoLevel:  logFilePath,
			logrus.WarnLevel:  logFilePath,
			logrus.ErrorLevel: logFilePath,
			logrus.FatalLevel: logFilePath,
			logrus.PanicLevel: logFilePath,
		},
		&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "time",
				logrus.FieldKeyLevel: "level",
				logrus.FieldKeyMsg:   "msg",
			},
			TimestampFormat: "2006-01-02T15:04:05-07:00", // Format: YYYY-MM-DDTHH:mm:ssZ
		},
	)

	// Add the lfHook to the logger
	log.AddHook(lfHook)
}

// GetLogger returns the logger instance
func GetLogger() *logrus.Logger {
	return log
}
