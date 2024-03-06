package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string  `json:"name,omitempty"`
	Age    int     `json:"age,omitempty"`
	Height float64 `json:"height,omitempty"`
}

func main() {
	// person 1
	person := Person{
		Name:   "Juan",
		Age:    30,
		Height: 1.75,
	}
	jsonData, err := json.Marshal(person)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println(string(jsonData))

	// person 2
	person2 := Person{
		Name:   "Pedro",
		Age:    25,
		Height: 2.0,
	}
	jsonData2, err := json.Marshal(person2)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println(string(jsonData2))
}
