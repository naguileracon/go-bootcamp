package service

import "one"

var (
	// ErrMsgNoStudents is the error message when there are no students
	ErrMsgNoStudents = "there are no students"
)

// NewStudentDefault creates a new student default service
func NewStudentDefault(rp one.StudentRepository) StudentDefault {
	return StudentDefault{rp}
}

// StudentDefault is the default student service
type StudentDefault struct {
	rp one.StudentRepository
}

func (s StudentDefault) GetStudentsDetails() (students []one.Student, err string) {
	// get all the students
	students = s.rp.Get()

	// check if there are no students
	if len(students) == 0 {
		err = ErrMsgNoStudents
		return
	}

	return
}
