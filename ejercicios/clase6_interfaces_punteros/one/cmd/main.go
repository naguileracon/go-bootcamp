package main

import (
	"fmt"
	"one"
	"one/repository"
	"one/service"
)

func main() {
	// dependencies
	var rp one.StudentRepository
	db := []one.Student{
		{ID: 1, StudentAttributes: one.StudentAttributes{
			Name:        "Nicolás",
			LastName:    "Aguilera",
			DateOfEntry: "19022024",
			DNI:         "123456789",
		}},
		{ID: 2, StudentAttributes: one.StudentAttributes{
			Name:        "Juan",
			LastName:    "Gonzales",
			DateOfEntry: "19022024",
			DNI:         "1234567891011",
		}},
		{ID: 3, StudentAttributes: one.StudentAttributes{
			Name:        "Laura",
			LastName:    "Fonseca",
			DateOfEntry: "19022024",
			DNI:         "12345678910",
		}},
	}
	rp = repository.NewStudentsSlice(db)
	sv := service.NewStudentDefault(rp)

	// use the service
	students, err := sv.GetStudentsDetails()
	if err != "" {
		println(err)
		return
	}

	fmt.Println(students)
}
