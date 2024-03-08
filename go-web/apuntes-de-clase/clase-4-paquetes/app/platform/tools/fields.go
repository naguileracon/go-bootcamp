package tools

import "errors"

var (
	ErrFileNotFound = errors.New("file not found")
)

func CheckFieldExists(field map[string]any, requiredFields ...string) (err error) {
	for _, field := range requiredFields {
		if _, ok := field[field]; !ok {
			err = ErrFileNotFound
			return
		}
	}
	return
}
