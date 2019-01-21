package logger

import (
	"encoding/json"
)

// LogEntry default log entry
type LogEntry struct {
	Message      string                 `json:"message"`
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

// SetKey sets the value for the custom key
func (e *LogEntry) SetKey(key string, value interface{}) {
	if e.Keys == nil {
		e.Keys = make(map[string]interface{})
	}
	e.Keys[key] = value
}
