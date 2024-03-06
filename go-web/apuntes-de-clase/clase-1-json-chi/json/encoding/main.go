package main

import (
	"encoding/json"
	"io"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// writer
	// ...
	var writter io.Writer = os.Stdout

	encoder := json.NewEncoder(writter)

	p := Person{
		Name: "John",
		Age:  30,
	}

	err := encoder.Encode(p)
	if err != nil {
		println(err.Error())
		return
	}
}
