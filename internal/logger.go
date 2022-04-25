package internal

import (
	"log"
	"os"
)

type Logger struct {
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
}

// Returns a new Logger
func NewLogger() *Logger {
	flags := log.LstdFlags

	return &Logger{
		Info:  log.New(os.Stderr, "INFO: ", flags),
		Warn:  log.New(os.Stderr, "WARNING: ", flags),
		Error: log.New(os.Stderr, "ERROR: ", flags),
	}
}
