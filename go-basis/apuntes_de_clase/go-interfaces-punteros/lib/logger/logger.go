package logger

type Logger interface {
	Info(format string, args ...any)
}
