package logger

import (
	"encoding/json"
	"time"
)

// LogEntry default log entry
type LogEntry struct {
	EventTime    string                 `json:"eventTime"`
	Message      string                 `json:"message"`
	SourceName   string                 `json:"sourceName"`
	ErrorCode    string                 `json:"errorCode,omitempty"`
	ErrorMessage string                 `json:"errorMessage,omitempty"`
	RequestID    string                 `json:"requestId"`
	LogLevel     string                 `json:"logLevel"`
	Keys         map[string]interface{} `json:"keys,omitempty"`
}

// Stringify marshal json to string
func (e *LogEntry) Stringify() string {
	raw, err := json.Marshal(*e)
	if err != nil {
		e.SetLogLevel("ERROR")
		return err.Error()
	}

	return string(raw)
}

// SetLogLevel sets the log level of the message
func (e *LogEntry) SetLogLevel(level string) { e.LogLevel = level }

// SetRequestID sets the request id
func (e *LogEntry) SetRequestID(id string) { e.RequestID = id }

// SetSourceName sets the source name in the log event
func (e *LogEntry) SetSourceName(name string) { e.SourceName = name }

// SetErrorCode sets the error code in the message if it is empty
func (e *LogEntry) SetErrorCode(code string) {
	if e.ErrorCode == "" {
		e.ErrorCode = code
	}
}

// GetMessage returns the message
func (e *LogEntry) GetMessage() string { return e.Message }

// SetEventTime sets the event time to the time in RFC3339
func (e *LogEntry) SetEventTime(t time.Time) {
	e.EventTime = t.Format(time.RFC3339Nano)
}

// SetKey sets the value for the custom key
func (e *LogEntry) SetKey(key string, value interface{}) {
	if e.Keys == nil {
		e.Keys = make(map[string]interface{})
	}
	e.Keys[key] = value
}
