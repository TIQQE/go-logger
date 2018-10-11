package logger

import (
	"fmt"
	"log"
)

type logger struct {
	id     string
	prefix string
}

var cwLogger logger

// Init initializes the logger with the request id and prefix
func Init(requestID string) {
	cwLogger = logger{
		id:     requestID,
		prefix: "[%s] {%s} ",
	}
}

func setPrefix(level string) {
	log.SetPrefix(fmt.Sprintf(cwLogger.prefix, level, cwLogger.id))
}

// InfoString logs a string message with INFO level
func InfoString(msg string) {
	Info(&LogEntry{Message: msg})
}

// Info a message with INFO level
func Info(msg ILogEntry) {
	setPrefix("INFO")
	log.Println(msg.Stringify())
}

// WarnString a string with WARNING level
func WarnString(msg string) {
	Warn(&LogEntry{Message: msg})
}

// Warn logs a message with WARNING level
func Warn(msg ILogEntry) {
	setPrefix("WARNING")
	log.Println(msg.Stringify())
}

// ErrorString logs a string with ERROR level
func ErrorString(msg string) {
	Error(&LogEntry{Message: msg})
}

// Error logs a msg with ERROR level
func Error(msg ILogEntry) {
	setPrefix("ERROR")
	log.Println(msg.Stringify())
}
