package repository

import (
	"one"
)

// NewStudentsSlice creates a new students slice structure
func NewStudentsSlice(students []one.Student) StudentsSlice {
	return StudentsSlice{students}
}

// StudentsSlice is a slice of students
type StudentsSlice struct {
	students []one.Student
}

// Get returns all the students
func (s StudentsSlice) Get() (st []one.Student) {
	st = make([]one.Student, len(s.students))
	copy(st, s.students)
	return
}
