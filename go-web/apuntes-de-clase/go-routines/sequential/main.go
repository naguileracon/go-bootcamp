package main

import "time"

func Do(n1, n2 int) {
	time.Sleep(250 * time.Millisecond)
	result := n1 + n2
	println(result)
}

func main() {
	start := time.Now()

	Do(1, 2)

	Do(3, 4)

	Do(5, 6)

	elapsed := time.Since(start)
	println(elapsed.Milliseconds())

}
