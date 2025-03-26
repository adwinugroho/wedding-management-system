package logger

import (
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

var globalLogger *logrus.Logger

func InitLogger() *logrus.Logger {
	globalLogger = logrus.New()
	formatter := &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", filepath.Base(f.File) + ":" + strconv.Itoa(f.Line)
		},
	}
	globalLogger.SetFormatter(formatter)
	globalLogger.SetReportCaller(true)
	globalLogger.SetLevel(logrus.InfoLevel)
	return globalLogger
}

func getCallerInfo() (string, string, int) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "", "", 0
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(loc).Format("2006-01-02 15:04:05"), filepath.Base(file), line
}

func LogInfo(message string) {
	time, _, _ := getCallerInfo()
	globalLogger.WithFields(logrus.Fields{
		"time": time,
	}).Info(message)
}

func LogWarn(message string) {
	time, _, _ := getCallerInfo()
	globalLogger.WithFields(logrus.Fields{
		"time": time,
	}).Warn(message)
}

func LogError(message string) {
	time, _, _ := getCallerInfo()
	globalLogger.WithFields(logrus.Fields{
		"time": time,
	}).Error(message)
}

func LogFatal(message string) {
	time, _, _ := getCallerInfo()
	globalLogger.WithFields(logrus.Fields{
		"time": time,
	}).Fatal(message)
}

func LogPanic(message string) {
	time, _, _ := getCallerInfo()
	globalLogger.WithFields(logrus.Fields{
		"time": time,
	}).Panic(message)
}

func LogWithFields(fields logrus.Fields, message string) {
	time, _, _ := getCallerInfo()
	fields["time"] = time
	globalLogger.WithFields(fields).Info(message)
}
