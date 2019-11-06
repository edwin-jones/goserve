package request

import (
	"fmt"

	"github.com/edwin-jones/goserve/status"
)

// Error A representation of an HTTP request error with error response data
type Error struct {
	message    string
	StatusCode status.Code
}

// newRequestError Constructor for RequestError
func newError(StatusCode status.Code) *Error {
	return &Error{
		message:    fmt.Sprintf("error %d", StatusCode),
		StatusCode: StatusCode,
	}
}

func (e *Error) Error() string {
	return e.message
}
