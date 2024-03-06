package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "customers.txt"
	file, err := os.Open(filename)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
		fmt.Println("Ejecución finalizada")
	}()

	defer func() {
		file.Close()
	}()

	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	fmt.Println("File opened successfully")

}
