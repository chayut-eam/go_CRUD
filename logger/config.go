package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
	"crud/chrono"
	"crud/model"
)

var (
	logger *log.Entry
)

func Init(appInfo model.AppInfo, logConfig model.LoggingConfig) {
	log.SetFormatter(logFormatter(logConfig))
	log.SetLevel(logLevel(logConfig))
	log.SetOutput(os.Stdout)
	logger = log.StandardLogger().WithField("application", appInfo.Name)
}

func Logger() *log.Entry {
	return logger
}

func logFormatter(config model.LoggingConfig) *log.JSONFormatter {
	return &log.JSONFormatter{
		TimestampFormat: string(chrono.Format_ISO8601),
	}
}

func logLevel(config model.LoggingConfig) log.Level {
	logLevel, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		panic(err)
	}
	return logLevel
}
