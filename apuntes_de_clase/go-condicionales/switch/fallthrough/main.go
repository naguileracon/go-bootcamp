package main

import "fmt"

func main() {

	var age int = 25

	switch {
	case age > 21:
		println("You can drink")
		fallthrough
	case age > 18:
		println("You can vote")
		fallthrough
	case age > 16:
		println("You can drive a car")
	default:
		println("not sure what you can do with that age")
	}

	number := 30
	switch {
	case number%15 == 0:
		fmt.Printf("%d is multiple of 15\n", number)
		fallthrough
	case number%5 == 0:
		fmt.Printf("%d is multiple of 5\n", number)
	default:
		fmt.Println("No condition met")
	}
}
