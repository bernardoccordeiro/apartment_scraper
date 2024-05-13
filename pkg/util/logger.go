package util

import (
	"io"
	"log"
)

// LogLevel type for specifying log levels
type LogLevel string

// Log level constants
const (
    LogLevelInfo LogLevel = "INFO"
    LogLevelWarning LogLevel = "WARNING"
    LogLevelError LogLevel = "ERROR"
    LogLevelFatal LogLevel = "FATAL"
)

// Logger wraps standard log.Logger from the Go standard library
type Logger struct {
    *log.Logger
}

// NewLogger creates a new Logger instance with a specified output and log level
func NewLogger(out io.Writer, level LogLevel) *Logger {
    logger := log.New(out, "", log.Ldate|log.Ltime|log.Lshortfile)
    switch level {
    case LogLevelInfo:
        logger.SetPrefix("INFO: ")
    case LogLevelWarning:
        logger.SetPrefix("WARNING: ")
    case LogLevelError:
        logger.SetPrefix("ERROR: ")
    case LogLevelFatal:
        logger.SetPrefix("FATAL: ")
    }
    return &Logger{logger}
}

func (lg *Logger) Write(p []byte) (int, error) {
    return lg.Writer().Write(p)
}
