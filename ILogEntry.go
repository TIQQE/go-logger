package logger

import "time"

// ILogEntry interface for log entries
type ILogEntry interface {
	Stringify() string
	SetLogLevel(string)
	SetRequestID(string)
	SetEventTime(time.Time)
	SetSourceName(string)
	SetErrorCode(string)
	GetMessage() string
}
