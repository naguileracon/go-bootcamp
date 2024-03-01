package main

import "fmt"

func main() {
	var temperature int = 17
	var atmosphericPressure float32 = 20.1
	var humidity float32 = 20.1

	fmt.Printf("%d %.1f %.1f\n", temperature, atmosphericPressure, humidity)
}
