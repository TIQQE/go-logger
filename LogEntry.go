package logger

import (
	"encoding/json"
)

type LogEntry struct {
	Message      string `json:"message"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (e *LogEntry) stringify() string {
	raw, err := json.Marshal(*e)
	if err != nil {
		setPrefix("ERROR")
		return err.Error()
	}

	return string(raw)
}
