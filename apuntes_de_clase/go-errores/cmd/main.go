package main

import (
	"errors"
	"fmt"
)

type CumstomErrorSalary struct {
	msg string
}

func (e *CumstomErrorSalary) Error() string {
	return e.msg
}

var (
	errorGlobal = "Error: salary is less than 10000"
	errorUp     = "Error: salary is much 10000"
)

func main() {
	salary := 9000

	err := salaryError(salary)

	errCumtom := CumstomErrorSalary{msg: errorGlobal}
	errUpCustom := CumstomErrorSalary{msg: errorUp}

	switch {
	case errors.Is(err, &errCumtom):
		fmt.Println("El salario es menor")
	case errors.Is(err, &errUpCustom):
		fmt.Println("El salario es mayor")
	default:
		fmt.Println("Nada")
	}
	fmt.Print("Hello")
}

func salaryError(salary int) error {
	if salary <= 10000 {
		return &CumstomErrorSalary{msg: errorGlobal}
	} else {
		return &CumstomErrorSalary{msg: errorUp}
	}
}
