package main

import "fmt"

var SalaryErrorMessage = "Error: the salary entered does not reach the taxable minimum"

func NewSalaryError() error {
	return &SalaryError{msg: SalaryErrorMessage}
}

type SalaryError struct {
	msg string
}

func (se *SalaryError) Error() string {
	return se.msg
}

func ValidateSalary(salary int) (err error) {
	if salary < 150000 {
		err = NewSalaryError()
		return
	}
	return
}
func main() {
	salary := 170000
	err := ValidateSalary(salary)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Must pay tax")
}
