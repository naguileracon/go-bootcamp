package repository

func NewVehicleRepositoryError(errorMessage string, httpStatusCode int) *VehicleRepositoryError {
	return &VehicleRepositoryError{ErrorMessage: errorMessage, HttpStatusCode: httpStatusCode}
}

type VehicleRepositoryError struct {
	ErrorMessage   string
	HttpStatusCode int
}

func (e *VehicleRepositoryError) Error() string {
	return e.ErrorMessage
}
