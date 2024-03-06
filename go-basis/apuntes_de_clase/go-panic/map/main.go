package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Nombre del archivo a leer
	filename := "customers.txt"

	_, err := ioutil.ReadFile(filename)

	defer func() {
		r := recover()
		println(r)
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Println("ejecución finalizada")
		}
	}()

	if err != nil {
		panic(nil)
	}

	fmt.Println("ejecución finalizada")
	// Imprimir "ejecución finalizada"
}
