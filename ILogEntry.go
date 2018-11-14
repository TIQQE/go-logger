package logger

// ILogEntry interface for log entries
type ILogEntry interface {
	Stringify() string
	SetLogLevel(string)
	SetRequestID(string)
}
