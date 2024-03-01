package main

import "fmt"

func main() {
	// default
	// var fruitsCounter map[string]int

	// initialize native
	fruitsCounter := map[string]int{
		"apple":  2,
		"banana": 3,
	}

	// with make
	fruitsCounter = make(map[string]int)
	fruitsCounter["apple"] = 2
	fruitsCounter["banana"] = 2
	fruitsCounter["orange"] = 2

	// check key exists
	mango, ok := fruitsCounter["banana"]
	if ok {
		fmt.Println("banana", mango)
	} else {
		fmt.Println("does not exist")

	}

}
