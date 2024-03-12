package validator

import (
	"app/internal"
)

// VehicleValidator is an interface that represents a validator for vehicles
type VehicleValidator interface {
	// Validate is a method that validates a product
	Validate(v *internal.VehicleAttributes) (err error)
}
