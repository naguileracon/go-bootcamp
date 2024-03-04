package main

import (
	"errors"
	"fmt"
)

var SalaryErrorMessage = errors.New("error: the salary entered does not reach the taxable minimum")

func ValidateSalary(salary int) (err error) {
	if salary < 10000 {
		err = SalaryErrorMessage
		return
	}
	return
}

func main() {
	salary := 10000
	err := ValidateSalary(salary)
	if errors.Is(err, SalaryErrorMessage) {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Must pay tax")
}
