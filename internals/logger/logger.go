package logger

import (
	"path/filepath"
	"runtime"

	"github.com/adwinugroho/wedding-management-system/internals/helpers"
	"github.com/sirupsen/logrus"
)

var globalLogger *logrus.Logger

func InitLogger() *logrus.Logger {
	globalLogger = logrus.New()
	globalLogger.SetFormatter(&logrus.JSONFormatter{})
	globalLogger.SetLevel(logrus.InfoLevel)
	return globalLogger
}

func getCallerInfo() (string, string, int) {
	_, file, line, ok := runtime.Caller(2) // 2 to skip the logger function itself
	if !ok {
		return "", "", 0
	}
	return helpers.TimeHostNow("Asia/Jakarta").Format("2006-01-02 15:04:05"), filepath.Base(file), line
}

func LogInfo(message string) {
	time, file, line := getCallerInfo()
	globalLogger.WithFields(logrus.Fields{
		"time": time,
		"file": file,
		"line": line,
	}).Info(message)
}

func LogWarn(message string) {
	time, file, line := getCallerInfo()
	globalLogger.WithFields(logrus.Fields{
		"time": time,
		"file": file,
		"line": line,
	}).Warn(message)
}

func LogError(message string) {
	time, file, line := getCallerInfo()
	globalLogger.WithFields(logrus.Fields{
		"time": time,
		"file": file,
		"line": line,
	}).Error(message)
}

func LogFatal(message string) {
	time, file, line := getCallerInfo()
	globalLogger.WithFields(logrus.Fields{
		"time": time,
		"file": file,
		"line": line,
	}).Fatal(message)
}

func LogPanic(message string) {
	time, file, line := getCallerInfo()
	globalLogger.WithFields(logrus.Fields{
		"time": time,
		"file": file,
		"line": line,
	}).Panic(message)
}

func LogWithFields(fields logrus.Fields, message string) {
	time, file, line := getCallerInfo()
	fields["time"] = time
	fields["file"] = file
	fields["line"] = line
	globalLogger.WithFields(fields).Info(message)
}
