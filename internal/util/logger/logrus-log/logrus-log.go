package logrus_log

import (
	"io"
	"runtime"

	"github.com/tadilbek11kz/ePharma-backend/internal/util/logger"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	logrus *logrus.Logger
}

func New() logger.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	return &LogrusLogger{
		log,
	}
}

func (logger *LogrusLogger) Info(message string, args map[string]interface{}) {
	fields := logrus.Fields{
		"method": checkFuncName(2),
	}
	for key, value := range args {
		fields[key] = value
	}

	logger.logrus.WithFields(fields).Info(message)
}

func (logger *LogrusLogger) Error(message string, args map[string]interface{}) {
	fields := logrus.Fields{
		"method": checkFuncName(2),
	}
	for key, value := range args {
		fields[key] = value
	}

	logger.logrus.WithFields(fields).Error(message)
}

func (logger *LogrusLogger) PanicLog(message string, args map[string]interface{}) {
	fields := logrus.Fields{
		"method": checkFuncName(5),
	}
	for key, value := range args {
		fields[key] = value
	}

	logger.logrus.WithFields(fields).Info(message)
	logger.logrus.SetOutput(io.Discard)
}

func checkFuncName(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		return "unknown"
	}
	me := runtime.FuncForPC(pc)
	if me == nil {
		return "unnamed"
	}
	return me.Name()
}
