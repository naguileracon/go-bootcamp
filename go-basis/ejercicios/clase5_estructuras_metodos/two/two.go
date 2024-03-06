package main

import "fmt"

func main() {
	person := Person{
		ID:          1,
		Name:        "Nicolas",
		DateOfBirth: "06101999",
	}
	employee := Employee{
		ID:       1,
		Position: "Trainee",
		Data:     person,
	}
	employee.PrintEmployee()
}

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Data     Person
}

func (employee Employee) PrintEmployee() {
	fmt.Printf("Employee id: %d \n", employee.ID)
	fmt.Printf("Employee position: %s \n", employee.Position)
	fmt.Printf("Employee data id: %d \n", employee.Data.ID)
	fmt.Printf("Employee name %s \n", employee.Data.Name)
	fmt.Printf("Employee date of birth %s \n", employee.Data.DateOfBirth)
}
