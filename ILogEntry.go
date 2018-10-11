package logger

type ILogEntry interface {
	stringify() string
}
