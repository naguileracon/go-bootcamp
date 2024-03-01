package main

import "fmt"

func main() {
	// slice initialize

	var fruits []string
	if fruits == nil {
		fmt.Println("names is nil")
	}

	//length = 3
	//capacity = 3
	var names []string = []string{
		"John",
		"Jane",
		"Joe",
	}
	fmt.Println(len(names), cap(names))

	// length = 0
	// capacity = 0
	var colors []string = []string{}
	fmt.Println(len(colors), cap(colors))

	// length = 0
	// capacity = 5
	var numbers []int = make([]int, 0, 5)
	fmt.Println(len(numbers), cap(numbers))

	numbers = append(numbers, 1, 2, 3, 4, 5)
	fmt.Println(len(numbers), cap(numbers))

	//access slice
	fmt.Println("numbers from position 0 to 4;", numbers[0:5])

}
