package main

import "time"

func Do(n1, n2 int, ch chan<- int) {
	time.Sleep(250 * time.Millisecond)
	result := n1 + n2
	ch <- result
}

func main() {
	start := time.Now()

	// channel with capacity 0
	// unbuffered channel
	ch := make(chan int, 3)

	go Do(1, 2, ch)

	go Do(3, 4, ch)

	go Do(5, 6, ch)

	// receive from channel
	//r1 := <-ch
	//r2 := <-ch
	//r3 := <-ch
	//println(r1, r2, r3)
	elapsed := time.Since(start)
	println(elapsed.Milliseconds())

}
