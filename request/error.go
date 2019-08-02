package request

import "fmt"

// ErrorCode HTTP error code typedef
type ErrorCode int

// HTTP error codes
const (
	BadRequest           ErrorCode = 400
	NotFound             ErrorCode = 404
	InvalidHTTPMethod    ErrorCode = 405
	URITooLong           ErrorCode = 414
	UnsupportedMediaType ErrorCode = 415
)

// Error A representation of an HTTP request error with error response data
type Error struct {
	message   string
	ErrorCode ErrorCode
}

// newRequestError Constructor for RequestError
func newError(errorCode ErrorCode) *Error {
	return &Error{
		message:   fmt.Sprintf("An invalid http request has been made: Error code %d", errorCode),
		ErrorCode: errorCode,
	}
}

func (e *Error) Error() string {
	return e.message
}
