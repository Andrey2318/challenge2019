package logger

import (
	"challenge2019/internal/config"
	"github.com/sirupsen/logrus"
	"os"

	"time"
)

var log *logrus.Logger

func Logger() *logrus.Logger {
	if log != nil {
		return log
	}

	log = logrus.New()
	log.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: false,
		TimestampFormat:  time.RFC3339,
		FieldMap:         nil,
		CallerPrettyfier: nil,
	}

	log.SetOutput(os.Stderr)

	switch config.GetLogLevel() {
	case "fatal":
		log.Level = logrus.FatalLevel
	case "error":
		log.Level = logrus.ErrorLevel
	case "warn":
		log.Level = logrus.WarnLevel
	case "info":
		log.Level = logrus.InfoLevel
	case "debug":
		log.Level = logrus.DebugLevel
	case "trace":
		log.Level = logrus.TraceLevel
	default:
		log.Level = logrus.PanicLevel
	}

	return log
}
