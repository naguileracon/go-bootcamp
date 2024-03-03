package one

// StudentAttributes is a struct that represents the attributes of a student
type StudentAttributes struct {
	Name        string
	LastName    string
	DNI         string
	DateOfEntry string
}

// Student is a struct that represents a student
type Student struct {
	// ID is the unique identifier of the student
	ID int
	// StudentsAttributes is the attributes of the student
	StudentAttributes
}
