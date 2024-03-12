package service

func NewVehicleServiceError(errorMessage string, httpStatusCode int) *VehicleServiceError {
	return &VehicleServiceError{ErrorMessage: errorMessage, HttpStatusCode: httpStatusCode}
}

type VehicleServiceError struct {
	ErrorMessage   string
	HttpStatusCode int
}

func (e *VehicleServiceError) Error() string {
	return e.ErrorMessage
}
