package main

func main() {
	var cc int
	var oddNumbers int

	for cc < 100 {
		cc += 3
		if cc%2 == 0 {
			continue
		}
		oddNumbers++
	}

	println("Odd numbers:", oddNumbers)
}
