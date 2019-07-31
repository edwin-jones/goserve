package parser

// RequestError A representation of an HTTP request error with error response data
type RequestError struct {
	message  string
	Response string
}

// newRequestError Constructor for RequestError
func newRequestError(response string) *RequestError {
	return &RequestError{
		message:  "An invalid http request has been made.",
		Response: response,
	}
}

func (e *RequestError) Error() string {
	return e.message
}
