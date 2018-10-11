package logger

import (
	"fmt"
	"log"
)

type logger struct {
	id     string
	prefix string
}

var cwLogger logger

// Init initializes the logger with the request id and prefix
func Init(requestID string) {
	cwLogger = logger{
		id:     requestID,
		prefix: "[%s] {%s} ",
	}
}

func setPrefix(level string) {
	log.SetPrefix(fmt.Sprintf(cwLogger.prefix, level, cwLogger.id))
}

func InfoString(msg string) {
	Info(&LogEntry{Message: msg})
}

func Info(msg ILogEntry) {
	setPrefix("INFO")
	log.Println(msg.stringify())
}

func WarnString(msg string) {
	Warn(&LogEntry{Message: msg})
}

func Warn(msg ILogEntry) {
	setPrefix("WARNING")
	log.Println(msg.stringify())
}

func ErrorString(msg string) {
	Error(&LogEntry{Message: msg})
}

func Error(msg ILogEntry) {
	setPrefix("ERROR")
	log.Println(msg.stringify())
}
