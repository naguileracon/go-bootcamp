package calculator

type MathOperation func(int, int) (int, string)

type MathOperator string

const (
	OperatorSum MathOperator = "+"
	OperatorSub MathOperator = "-"
	OperatorMul MathOperator = "*"
	OperatorDiv MathOperator = "/"
)

var (
	ErrMsgDivisionIndeterminate string = "error division indeterminate"
	ErrorDivisionBy0            string = "error by 0"
)

// Add calculates the sum of two numbers
// these numbers are integers
func Add(a, b int) (result int, err string) {
	result = a + b
	return
}

func Substract(a, b int) (int, string) {
	return a - b, ""
}

func Multiply(a, b int) (int, string) {
	return a * b, ""
}

func Divide(a, b int) (result int, err string) {
	if a == 0 && b == 0 {
		err = ErrMsgDivisionIndeterminate
		return
	}

	if b == 0 {
		err = ErrorDivisionBy0
		return
	}
	result = a / b
	return
}
