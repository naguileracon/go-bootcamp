package one

// StudentRepository is an interface that represents the student repository
type StudentRepository interface {
	// Get returns all the students
	Get() (st []Student)
}
