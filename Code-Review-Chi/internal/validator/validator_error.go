package validator

func NewVehicleValidatorError(errorMessage string, httpStatusCode int) *VehicleValidatorError {
	return &VehicleValidatorError{errorMessage: errorMessage, httpStatusCode: httpStatusCode}
}

type VehicleValidatorError struct {
	errorMessage   string
	httpStatusCode int
}

func (e *VehicleValidatorError) Error() string {
	return e.errorMessage
}
