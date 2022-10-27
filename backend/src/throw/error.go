package throw

import (
	"fmt"
)

type HttpError struct {
	message    string
	statusCode int
}

func NewHttpError(message string, statusCode int) *HttpError {
	return &HttpError{
		message,
		statusCode,
	}
}

func (I *HttpError) GetMessage() string {
	return I.message
}

func (I *HttpError) GetStatusCode() int {
	return I.statusCode
}

type ServerError struct {
	err       error
	httpError *HttpError
}

func NewServerError(err error, httpError *HttpError) *ServerError {
	if err == nil {
		return &ServerError{
			err:       fmt.Errorf("error: %s", httpError.GetMessage()),
			httpError: httpError,
		}
	}

	return &ServerError{
		err,
		httpError,
	}
}

func (I *ServerError) GetError() error {
	return I.err
}

func (I *ServerError) Error() string {
	return I.GetError().Error()
}

func (I *ServerError) GetHttpError() *HttpError {
	return I.httpError
}
