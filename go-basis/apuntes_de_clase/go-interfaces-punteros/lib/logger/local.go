package logger

import (
	"github.com/fatih/color"
)

func NewLocal(color string) Logger {
	return Local{color: color}
}

type Local struct {
	color string
}

func (l Local) Info(format string, args ...any) {
	switch l.color {
	case "red":
		color.Red(format, args...)
	case "green":
		color.Green(format, args...)
	default:
		color.White(format, args...)
	}
}
