package logger

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"testing"
)

func TestLogging(t *testing.T) {
	Init("test-req", "TEST")

	t.Run("Info()", testInfo)
	t.Run("InfoString()", testInfoString)
	t.Run("InfoStringf()", testInfoStringf)

	t.Run("Warn()", testWarn)
	t.Run("WarnString()", testWarnString)
	t.Run("WarnStringf()", testWarnStringf)

	t.Run("Error()", testError)
	t.Run("ErrorString()", testErrorString)
	t.Run("ErrorStringf()", testErrorStringf)
}

func testInfoString(t *testing.T) {
	testCases := []struct {
		input    string
		expected LogEntry
	}{
		{"Test", LogEntry{Message: "Test", SourceName: "TEST", RequestID: "test-req", LogLevel: "INFO"}},
		{"Longer Test", LogEntry{Message: "Longer Test", SourceName: "TEST", RequestID: "test-req", LogLevel: "INFO"}},
	}

	for i, tc := range testCases {
		output := captureOutput(t, func() {
			InfoString(tc.input)
		})

		entry := unmarshal(t, output)

		if !evalEntry(t, entry, tc.expected) {
			t.Logf("Test[%d]: %s", i, output)
		}

	}
}

func testWarnString(t *testing.T) {
	testCases := []struct {
		input    string
		expected LogEntry
	}{
		{"Test", LogEntry{Message: "Test", ErrorCode: "Test", SourceName: "TEST", RequestID: "test-req", LogLevel: "WARNING"}},
		{"Longer Test", LogEntry{Message: "Longer Test", ErrorCode: "Longer Test", SourceName: "TEST", RequestID: "test-req", LogLevel: "WARNING"}},
	}

	for i, tc := range testCases {
		output := captureOutput(t, func() {
			WarnString(tc.input)
		})

		entry := unmarshal(t, output)

		if !evalEntry(t, entry, tc.expected) {
			t.Logf("Test[%d]: %s", i, output)
		}

	}
}

func testErrorString(t *testing.T) {
	testCases := []struct {
		input    string
		expected LogEntry
	}{
		{"Test", LogEntry{Message: "Test", ErrorCode: "Test", SourceName: "TEST", RequestID: "test-req", LogLevel: "ERROR", Action: "Open"}},
		{"Longer Test", LogEntry{Message: "Longer Test", ErrorCode: "Longer Test", SourceName: "TEST", RequestID: "test-req", LogLevel: "ERROR", Action: "Open"}},
	}

	for i, tc := range testCases {
		output := captureOutput(t, func() {
			ErrorString(tc.input)
		})

		entry := unmarshal(t, output)

		if !evalEntry(t, entry, tc.expected) {
			t.Logf("Test[%d]: %s", i, output)
		}

	}
}

func testInfoStringf(t *testing.T) {
	testCases := []struct {
		input    string
		args     []interface{}
		expected LogEntry
	}{
		{"Test %d", []interface{}{1}, LogEntry{Message: "Test 1", SourceName: "TEST", RequestID: "test-req", LogLevel: "INFO"}},
		{"Complex %s(%T)", []interface{}{"Test", "Test"}, LogEntry{Message: "Complex Test(string)", SourceName: "TEST", RequestID: "test-req", LogLevel: "INFO"}},
	}

	for i, tc := range testCases {
		output := captureOutput(t, func() {
			InfoStringf(tc.input, tc.args...)
		})

		entry := unmarshal(t, output)

		if !evalEntry(t, entry, tc.expected) {
			t.Logf("Test[%d]: %s", i, output)
		}

	}
}

func testWarnStringf(t *testing.T) {
	testCases := []struct {
		input    string
		args     []interface{}
		expected LogEntry
	}{
		{"Test %d", []interface{}{1}, LogEntry{Message: "Test 1", ErrorCode: "Test 1", SourceName: "TEST", RequestID: "test-req", LogLevel: "WARNING"}},
		{"Complex %s(%T)", []interface{}{"Test", "Test"}, LogEntry{Message: "Complex Test(string)", ErrorCode: "Complex Test(string)", SourceName: "TEST", RequestID: "test-req", LogLevel: "WARNING"}},
	}

	for i, tc := range testCases {
		output := captureOutput(t, func() {
			WarnStringf(tc.input, tc.args...)
		})

		entry := unmarshal(t, output)

		if !evalEntry(t, entry, tc.expected) {
			t.Logf("Test[%d]: %s", i, output)
		}

	}
}

func testErrorStringf(t *testing.T) {
	testCases := []struct {
		input    string
		args     []interface{}
		expected LogEntry
	}{
		{"Test %d", []interface{}{1}, LogEntry{Message: "Test 1", ErrorCode: "Test 1", SourceName: "TEST", RequestID: "test-req", LogLevel: "ERROR", Action: "Open"}},
		{"Complex %s(%T)", []interface{}{"Test", "Test"}, LogEntry{Message: "Complex Test(string)", ErrorCode: "Complex Test(string)", SourceName: "TEST", RequestID: "test-req", LogLevel: "ERROR", Action: "Open"}},
	}

	for i, tc := range testCases {
		output := captureOutput(t, func() {
			ErrorStringf(tc.input, tc.args...)
		})

		entry := unmarshal(t, output)

		if !evalEntry(t, entry, tc.expected) {
			t.Logf("Test[%d]: %s", i, output)
		}

	}
}

func testInfo(t *testing.T) {
	testCases := []struct {
		input    LogEntry
		expected LogEntry
	}{
		{LogEntry{Message: "Test"}, LogEntry{Message: "Test", SourceName: "TEST", RequestID: "test-req", LogLevel: "INFO"}},
		{
			LogEntry{
				Message: "This is a more complex test",
				Keys: map[string]interface{}{
					"Key1": "also",
					"Key2": "This is a key",
				},
			},
			LogEntry{
				Message: "This is a more complex test",
				Keys: map[string]interface{}{
					"Key1": "also",
					"Key2": "This is a key",
				},
				SourceName: "TEST",
				RequestID:  "test-req",
				LogLevel:   "INFO",
			},
		},
	}

	for i, tc := range testCases {
		output := captureOutput(t, func() {
			Info(&tc.input)
		})

		entry := unmarshal(t, output)

		if !evalEntry(t, entry, tc.expected) {
			t.Logf("Test[%d]: %s", i, output)
		}

	}
}

func testWarn(t *testing.T) {
	testCases := []struct {
		input    LogEntry
		expected LogEntry
	}{
		{LogEntry{Message: "Test", ErrorMessage: "Warning"}, LogEntry{Message: "Test", ErrorMessage: "Warning", ErrorCode: "Test", SourceName: "TEST", RequestID: "test-req", LogLevel: "WARNING"}},
		{
			LogEntry{
				Message: "This is a more complex test",
				Keys: map[string]interface{}{
					"Key1": "also",
					"Key2": "This is a key",
				},
			},
			LogEntry{
				Message:      "This is a more complex test",
				ErrorMessage: "",
				ErrorCode:    "This is a more complex test",
				Keys: map[string]interface{}{
					"Key1": "also",
					"Key2": "This is a key",
				},
				SourceName: "TEST",
				RequestID:  "test-req",
				LogLevel:   "WARNING",
			},
		},
	}

	for i, tc := range testCases {
		output := captureOutput(t, func() {
			Warn(&tc.input)
		})

		entry := unmarshal(t, output)

		if !evalEntry(t, entry, tc.expected) {
			t.Logf("Test[%d]: %s", i, output)
		}

	}
}

func testError(t *testing.T) {
	testCases := []struct {
		input    LogEntry
		expected LogEntry
	}{
		{LogEntry{Message: "Test", ErrorMessage: "Warning"}, LogEntry{Message: "Test", ErrorMessage: "Warning", ErrorCode: "Test", SourceName: "TEST", RequestID: "test-req", LogLevel: "ERROR", Action: "Open"}},
		{
			LogEntry{
				Message: "This is a more complex test",
				Keys: map[string]interface{}{
					"Key1": "also",
					"Key2": "This is a key",
				},
			},
			LogEntry{
				Message:      "This is a more complex test",
				ErrorMessage: "",
				ErrorCode:    "This is a more complex test",
				Keys: map[string]interface{}{
					"Key1": "also",
					"Key2": "This is a key",
				},
				SourceName: "TEST",
				RequestID:  "test-req",
				LogLevel:   "ERROR",
				Action:     "Open",
			},
		},
	}

	for i, tc := range testCases {
		output := captureOutput(t, func() {
			Error(&tc.input)
		})

		entry := unmarshal(t, output)

		if !evalEntry(t, entry, tc.expected) {
			t.Logf("Test[%d]: %s", i, output)
		}

	}
}

func evalEntry(t *testing.T, entry, expected LogEntry) bool {
	if entry.Message != expected.Message {
		t.Errorf("Message is wrong. got=%s want=%s", entry.Message, expected.Message)
		return false
	}

	if entry.RequestID != expected.RequestID {
		t.Errorf("RequestID wrong. got=%s want=%s", entry.RequestID, expected.RequestID)
		return false
	}

	if entry.LogLevel != expected.LogLevel {
		t.Errorf("LogLevel wrong. got=%s want=%s", entry.LogLevel, expected.LogLevel)
		return false
	}

	if entry.SourceName != expected.SourceName {
		t.Errorf("SourceName is wrong. got=%s want=%s", entry.SourceName, expected.SourceName)
		return false
	}

	if entry.ErrorCode != expected.ErrorCode {
		t.Errorf("ErrorCode is wrong. got=%s want=%s", entry.ErrorCode, expected.ErrorCode)
		return false
	}

	if entry.Action != expected.Action {
		t.Errorf("Action is wrong. got=%s want=%s", entry.Action, expected.Action)
		return false
	}

	if entry.ErrorMessage != expected.ErrorMessage {
		t.Errorf("ErrorMessage is wrong. got=%s want=%s", entry.ErrorMessage, expected.ErrorMessage)
		return false
	}

	if entry.Keys == nil && expected.Keys != nil {
		t.Errorf("Keys was expected but not found")
		return false
	} else if entry.Keys != nil && expected.Keys == nil {
		t.Errorf("Expected no Keys")
		return false
	} else if entry.Keys != nil && expected.Keys != nil {
		numKeys := len(entry.Keys)
		expectedNum := len(expected.Keys)

		if numKeys != expectedNum {
			t.Errorf("Number of keys is wrong. got=%d want=%d", numKeys, expectedNum)
			return false
		}

		for k, v := range expected.Keys {
			val, ok := entry.Keys[k]
			if !ok {
				t.Errorf("%s is missing", k)
				return false
			}

			if val != v {
				t.Errorf("Value of %s is wrong. got=%v(%T) want=%v(%T)", k, val, val, v, v)
				return false
			}

		}
	}
	return true
}

func unmarshal(t *testing.T, str string) (entry LogEntry) {
	if err := json.Unmarshal([]byte(str), &entry); err != nil {
		t.Errorf("failed to unmarshal log event: %v", err)
	}
	return
}

func captureOutput(t *testing.T, f func()) string {

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe for stdoutput: %v", err)
	}

	old := os.Stdout
	os.Stdout = w
	out := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		out <- buf.String()
	}()
	f()

	w.Close()
	os.Stdout = old

	return <-out
}
