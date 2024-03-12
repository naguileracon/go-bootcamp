package platform

import (
	"fmt"
	"net/http"
)

// ValidateRequiredFields is a method that validates required fields for a vehicle
func ValidateRequiredFields(fields map[string]any, requiredFields ...string) (err *PlatformError) {
	for _, field := range requiredFields {
		if _, ok := fields[field]; !ok {
			err = NewPlatformError(fmt.Sprintf("Missing required field: %s", field), http.StatusBadRequest)
			return
		}
	}
	return
}
