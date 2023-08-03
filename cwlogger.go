package logger

import (
	"fmt"
	"time"
)

type logger struct {
	id           string
	sourceName   string
	sourceID     string
	debugEnabled bool

	sessionValues map[string]interface{}
}

var cwLogger logger

func (l logger) mergeSessionValues(entry interface{}) {
	e, ok := entry.(KeyValueHolder)
	if !ok {
		return
	}

	for k, v := range l.sessionValues {
		e.SetKey(k, v)
	}
}

// Init initializes the logger with the request id and prefix. Using this initialization method will disable debug logging.
func Init(requestID, sourceName string) {
	InitWithDebugLevel(requestID, sourceName, false)
}

// InitWithDebugLevel initializes the logger with the request id and prefix. The parameter debugEnabled specifies if logging of log level Debug will be enabled or not.
func InitWithDebugLevel(requestID, sourceName string, debugEnabled bool) {
	cwLogger = logger{
		id:            requestID,
		sourceName:    sourceName,
		debugEnabled:  debugEnabled,
		sessionValues: map[string]interface{}{},
	}
}

// WithKeysValue adds the value to the session map that will be merged into all log entries.
func WithKeysValue(key string, value interface{}) {
	cwLogger.sessionValues[key] = value
}

// DebugStringf debug log helper to use sprintf formatting.
func DebugStringf(format string, args ...interface{}) {
	if !cwLogger.debugEnabled {
		return
	}
	DebugString(fmt.Sprintf(format, args...))
}

// DebugString logs a string message with DEBUG level
func DebugString(msg string) {
	if !cwLogger.debugEnabled {
		return
	}
	Debug(&LogEntry{Message: msg})
}

// Debug logs a message with DEBUG level
func Debug(msg ILogEntry) {
	if !cwLogger.debugEnabled {
		return
	}
	msg.SetLogLevel("DEBUG")
	msg.SetRequestID(cwLogger.id)
	msg.SetEventTime(time.Now())
	msg.SetSourceName(cwLogger.sourceName)
	cwLogger.mergeSessionValues(msg)
	fmt.Println(msg.Stringify())
}

// InfoStringf info log helper to use sprintf formatting.
func InfoStringf(format string, args ...interface{}) {
	InfoString(fmt.Sprintf(format, args...))
}

// InfoString logs a string message with INFO level
func InfoString(msg string) {
	Info(&LogEntry{Message: msg})
}

// Info logs a message with INFO level
func Info(msg ILogEntry) {
	msg.SetLogLevel("INFO")
	msg.SetRequestID(cwLogger.id)
	msg.SetEventTime(time.Now())
	msg.SetSourceName(cwLogger.sourceName)
	cwLogger.mergeSessionValues(msg)
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
	cwLogger.mergeSessionValues(msg)
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
	msg.SetAction(ActionOpen)
	cwLogger.mergeSessionValues(msg)
	fmt.Println(msg.Stringify())
}
