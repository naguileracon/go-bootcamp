package logger

import (
	"fmt"
	"os"
)

func NewTextFile(path string) Logger {
	return TxtFile{Path: path}
}

type TxtFile struct {
	Path string
}

func (t TxtFile) Info(format string, args ...any) {
	file, err := os.OpenFile(t.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.WriteString(fmt.Sprintf(format, args...))
}
