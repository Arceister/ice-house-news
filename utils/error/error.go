package utils

type ErrorMessage struct {
	StatusCode    int
	ErrorResponse string
	ErrorMessage  string
}

type IErrorMessage interface {
	Message() string
	Status() int
	Error() string
}

func (e *ErrorMessage) Error() string {
	return e.ErrorResponse
}

func (e *ErrorMessage) Status() int {
	return e.StatusCode
}

func (e *ErrorMessage) Message() string {
	return e.ErrorMessage
}
