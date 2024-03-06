package main

func main() {
	cont := 0
	limit := 10
	for cont < limit {
		cont++
		println(cont)
	}
	println("cont is 10")

	for i := 0; i <= 20; i += 5 {
		println("counter", i)
	}
}
