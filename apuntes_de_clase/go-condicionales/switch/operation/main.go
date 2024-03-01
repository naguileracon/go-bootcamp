package main

func main() {
	var number int = 10

	switch r := number % 2; r {
	case 0:
		println("Even")
	default:
		println("Odd")
	}
}
