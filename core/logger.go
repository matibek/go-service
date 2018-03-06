package core

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Gets default application logger
func getLogger() *logrus.Logger {
	formatter := &logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 03:04:05",
	}
	log := logrus.New()
	log.Formatter = formatter
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)
	return log
}

// Customize logger from config
func initLogger(config map[string]string) {
	if level, ok := config["level"]; ok {
		lvl, err := logrus.ParseLevel(level)
		if err == nil {
			Logger.SetLevel(lvl)
		}
	}
	if filename, ok := config["file"]; ok {
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			panic("Failed to log to file!")
		}
		Logger.Out = file
	}
	if formatter, ok := config["formatter"]; ok && formatter == "json" {
		Logger.Formatter = &logrus.JSONFormatter{}
	}
}
