package logger

import (
	"fmt"
	"time"
)

type logger struct {
	id         string
	sourceName string
	sourceID   string
}

var cwLogger logger

// Init initializes the logger with the request id and prefix
func Init(requestID, sourceName string) {
	cwLogger = logger{
		id:         requestID,
		sourceName: sourceName,
	}
}

// InfoStringf info log helper to use sprintf formatting.
func InfoStringf(format string, args ...interface{}) {
	InfoString(fmt.Sprintf(format, args...))
}

// InfoString logs a string message with INFO level
func InfoString(msg string) {
	Info(&LogEntry{Message: msg})
}

// Info a message with INFO level
func Info(msg ILogEntry) {
	msg.SetLogLevel("INFO")
	msg.SetRequestID(cwLogger.id)
	msg.SetEventTime(time.Now())
	msg.SetSourceName(cwLogger.sourceName)
	fmt.Println(msg.Stringify())
}

// WarnStringf warn log helper to use sprintf formatting.
func WarnStringf(format string, args ...interface{}) {
	WarnString(fmt.Sprintf(format, args...))
}

// WarnString a string with WARNING level
func WarnString(msg string) {
	Warn(&LogEntry{Message: msg})
}

// Warn logs a message with WARNING level
func Warn(msg ILogEntry) {
	msg.SetLogLevel("WARNING")
	msg.SetRequestID(cwLogger.id)
	msg.SetEventTime(time.Now())
	msg.SetSourceName(cwLogger.sourceName)
	msg.SetErrorCode(msg.GetMessage())
	fmt.Println(msg.Stringify())
}

// ErrorStringf error log helper to use sprintf formatting.
func ErrorStringf(format string, args ...interface{}) {
	ErrorString(fmt.Sprintf(format, args...))
}

// ErrorString logs a string with ERROR level
func ErrorString(msg string) {
	Error(&LogEntry{Message: msg})
}

// Error logs a msg with ERROR level
func Error(msg ILogEntry) {
	msg.SetLogLevel("ERROR")
	msg.SetRequestID(cwLogger.id)
	msg.SetEventTime(time.Now())
	msg.SetSourceName(cwLogger.sourceName)
	msg.SetErrorCode(msg.GetMessage())
	fmt.Println(msg.Stringify())
}
