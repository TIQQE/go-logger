package logger

import (
	"encoding/json"
)

// LogEntry default log entry
type LogEntry struct {
	Message      string `json:"message"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Stringify marshal json to string
func (e *LogEntry) Stringify() string {
	raw, err := json.Marshal(*e)
	if err != nil {
		setPrefix("ERROR")
		return err.Error()
	}

	return string(raw)
}
