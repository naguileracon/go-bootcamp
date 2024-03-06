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
	if salary < 150000 {
		err = SalaryMinimumError
		err = fmt.Errorf("error: the minimum taxable amount is 150,000 and the salary entered is %d, %w", salary, err)
		return
	}
	return
}

func main() {
	salary := 200000
	err := ValidateSalary(salary)
	var salaryError *SalaryError
	if errors.As(err, &salaryError) {
		if salaryError.msg == "Error: the salary entered does not reach the taxable minimum" {
			fmt.Println(err.Error())
			return
		}
	}
	fmt.Println("Must pay tax")
}
