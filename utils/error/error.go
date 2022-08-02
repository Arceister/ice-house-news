package utils

type ErrorMessage struct {
	StatusCode   int
	ErrorMessage string
}

type IErrorMessage interface {
	Message() string
	Status() int
}

func (e *ErrorMessage) Status() int {
	return e.StatusCode
}

func (e *ErrorMessage) Message() string {
	return e.ErrorMessage
}
