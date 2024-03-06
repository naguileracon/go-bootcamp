package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	rd := strings.NewReader(`{"name": "John", "age": 30}`)

	decoder := json.NewDecoder(rd)

	var p Person
	err := decoder.Decode(&p)
	if err != nil {
		println(err.Error())
		return
	}

	fmt.Println(p)
}
