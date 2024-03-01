package main

import "fmt"

func main() {

	songs := []string{
		"imagine",
		"hotel california",
		"bohemian rhapsody",
		"imagine",
		"imagine",
		"imagine",
	}

	for ix, song := range songs {
		fmt.Println(ix, song)
	}

	// check repeated songs

	frequency := make(map[string]int)
	for _, song := range songs {
		//if _, ok := frequency[song]; ok {
		//	frequency[song]++
		//} else {
		//	frequency[song] = 1
		//}
		frequency[song]++
	}
	fmt.Println(frequency)
}
