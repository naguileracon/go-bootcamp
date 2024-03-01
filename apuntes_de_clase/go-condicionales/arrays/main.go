package main

import "fmt"

func main() {
	var names [3]string
	fmt.Println(names)

	var temps [5]float64 = [5]float64{
		1.2,
		1.3,
		1.4,
		1.5,
	}
	temps[4] = 1.6
	fmt.Println(temps)
}
