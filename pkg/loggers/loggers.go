package loggers

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
	logrusadapter "logur.dev/adapter/logrus"
	"logur.dev/logur"
)

func NewLogger() *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	return log
}

func NewLoggerTemporal(l *logrus.Logger) logur.KVLoggerFacade {
	return logur.LoggerToKV(logrusadapter.New(l))

}
