package logger

import (
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

var globalLogger *logrus.Logger

func InitLogger() *logrus.Logger {
	globalLogger = logrus.New()
	globalLogger.SetFormatter(&logrus.JSONFormatter{})
	globalLogger.SetLevel(logrus.InfoLevel)
	return globalLogger
}

func getCallerInfo() (string, int) {
	_, file, line, ok := runtime.Caller(2) // 2 to skip the logger function itself
	if !ok {
		return "unknown", 0
	}
	return filepath.Base(file), line
}

func LogInfo(message string) {
	file, line := getCallerInfo()
	globalLogger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Info(message)
}

func LogWarn(message string) {
	file, line := getCallerInfo()
	globalLogger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Warn(message)
}

func LogError(message string) {
	file, line := getCallerInfo()
	globalLogger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Error(message)
}

func LogFatal(message string) {
	file, line := getCallerInfo()
	globalLogger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Fatal(message)
}

func LogPanic(message string) {
	file, line := getCallerInfo()
	globalLogger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Panic(message)
}

func LogWithFields(fields logrus.Fields, message string) {
	file, line := getCallerInfo()
	fields["file"] = file
	fields["line"] = line
	globalLogger.WithFields(fields).Info(message)
}
