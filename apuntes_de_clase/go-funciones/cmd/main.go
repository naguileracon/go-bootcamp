package main

import "go-funciones/internal/calculator"

func main() {
	a := 10
	b := 5

	fnDiv, err := calculator.Orchestrator(calculator.OperatorDiv)

	if err != "" {
		println(err)
		return
	}

	result, err := fnDiv(a, b)
	if err != "" {
		println(err)
		return
	}
	println(result)
}
