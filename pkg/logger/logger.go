package logger

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "REVERSE-PROXY: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Info(msg string) {
	l.Printf("INFO: %s", msg)
}

func (l *Logger) Error(err error, msg string) {
	l.Printf("ERROR: %s - %v", msg, err)
}
