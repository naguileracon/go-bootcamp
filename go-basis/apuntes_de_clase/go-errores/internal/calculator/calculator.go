package calculator

import (
	"errors"
	"fmt"
)

var (
	ErrMsgDivisionIndeterminate = errors.New("error division indeterminate")
	ErrorDivisionBy0            = errors.New("error by 0")
)

func Divide(a, b int) (result int, err error) {
	if a == 0 && b == 0 {
		err = New(1000, a, b)
		err = fmt.Errorf("indeterminate division %w", err)
		return
	}

	if b == 0 {
		err = New(1001, a, b)
		err = fmt.Errorf("division by 0 %w", err)
		return
	}
	result = a / b
	return
}
