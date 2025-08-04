package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func InitLogger(logLevel string, logPath string) {
	log = logrus.New()
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.Fatal("Can not understand logLevel", err)
	}
	log.SetLevel(level)
	log.SetFormatter(&logrus.JSONFormatter{})
 
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file", err)
	}
	log.SetOutput(io.MultiWriter(file, os.Stdout))
}

func Debug(message string, fields map[string]interface{}) {
    log.WithFields(fields).Debug(message)
}

// Info 记录 Info 级别日志
func Info(message string, fields map[string]interface{}) {
    log.WithFields(fields).Info(message)
}

// Warn 记录 Warn 级别日志
func Warn(message string, fields map[string]interface{}) {
    log.WithFields(fields).Warn(message)
}

// Error 记录 Error 级别日志
func Error(message string, fields map[string]interface{}) {
    log.WithFields(fields).Error(message)
}

// Fatal 记录 Fatal 级别日志
func Fatal(message string, fields map[string]interface{}) {
    log.WithFields(fields).Fatal(message)
}