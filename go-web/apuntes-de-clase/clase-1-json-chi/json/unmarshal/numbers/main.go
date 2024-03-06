package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// person 1
	jsondata := []byte(`{
		"name":   "Juan",
		"age":    30,
		"height": 1.75,
		"weight": 80.0
	}`)

	m := make(map[string]any)
	err := json.Unmarshal(jsondata, &m)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	age, ok := m["age"].(float64)
	if !ok {
		fmt.Println("age is not an int")
		return
	}

	ageInt := int(age)
	ageInt++
	println(ageInt)
}
