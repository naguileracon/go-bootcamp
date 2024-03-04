package calculator

import "fmt"

func New(code int, args []any) error {
	return &MathError{Code: code, Args: args}
}

type MathError struct {
	Code int
	Args []any
}

func (e *MathError) Error() string {
	return fmt.Sprintf("math error: code %d args %v", e.Code, e.Args)
}
