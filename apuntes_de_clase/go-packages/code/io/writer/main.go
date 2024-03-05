package main

import (
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout

	bytes := []byte("Hello, Writer! How are you?")
	n, err := w.Write(bytes)
	if err != nil {
		println(err.Error())
		return
	}

	println("\n", n, "bytes written")

}
