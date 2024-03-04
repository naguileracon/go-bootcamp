package main

import (
	"errors"
	"fmt"
)

var SalaryMinimumError = NewSalaryError("Error: the salary entered does not reach the taxable minimum")

func NewSalaryError(message string) error {
	return &SalaryError{msg: message}
}

type SalaryError struct {
	msg string
}

func (se *SalaryError) Error() string {
	return se.msg
}

func ValidateSalary(salary int) (err error) {
	if salary < 10000 {
		err = SalaryMinimumError
		return
	}
	return
}

func main() {
	salary := 5000
	err := ValidateSalary(salary)
	if errors.Is(err, SalaryMinimumError) {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Must pay tax")
}
