package request

import "fmt"

// StatusCode HTTP status code typedef
type StatusCode int

// HTTP status codes
const (
	Success              StatusCode = 200
	BadRequest           StatusCode = 400
	NotFound             StatusCode = 404
	InvalidHTTPMethod    StatusCode = 405
	URITooLong           StatusCode = 414
	UnsupportedMediaType StatusCode = 415
)

// Error A representation of an HTTP request error with error response data
type Error struct {
	message    string
	StatusCode StatusCode
}

// newRequestError Constructor for RequestError
func newError(StatusCode StatusCode) *Error {
	return &Error{
		message:    fmt.Sprintf("An invalid http request has been made: Status code %d", StatusCode),
		StatusCode: StatusCode,
	}
}

func (e *Error) Error() string {
	return e.message
}
