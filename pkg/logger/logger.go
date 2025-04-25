package logger

import (
	"log"
)

type Logger struct {
	logger *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		logger: log.New(log.Writer(), "[FileCompressor] ", log.LstdFlags),
	}
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.logger.Printf("INFO: "+format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.logger.Printf("ERROR: "+format, v...)
}
