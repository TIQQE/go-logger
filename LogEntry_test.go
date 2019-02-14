package logger

import (
	"testing"
	"time"
)

func TestEventTime(t *testing.T) {
	eventTime, _ := time.Parse("2006-01-02T15:04:05.000", "2019-02-12T18:03:12.123")

	entry := LogEntry{}
	entry.SetEventTime(eventTime)
	expected := "2019-02-12T18:03:12.123Z"

	if entry.EventTime != expected {
		t.Errorf("EventTime wrong. expected=%s got=%s", expected, entry.EventTime)
	}

}

func TestAction(t *testing.T) {

	Init("test-req-id", "TestAction()")
	entry := LogEntry{
		Message: "Test",
	}

	Error(&entry)
	if entry.Action != ActionOpen {
		t.Errorf("Action wrong. got=%s expected=%s", entry.Action, ActionOpen)
	}

}
