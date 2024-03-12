package platform

func NewPlatformError(errorMessage string, httpStatusCode int) *PlatformError {
	return &PlatformError{ErrorMessage: errorMessage, HttpStatusCode: httpStatusCode}
}

type PlatformError struct {
	ErrorMessage   string
	HttpStatusCode int
}

func (e *PlatformError) Error() string {
	return e.ErrorMessage
}
