package main

// HTTPRequestError A representation of an HTTP error with error response data
type HTTPRequestError struct {
	message  string
	response []byte
}

// NewHTTPRequestError Constructor
func NewHTTPRequestError(response []byte) *HTTPRequestError {
	return &HTTPRequestError{
		message:  "An invalid http request has been made.",
		response: response,
	}
}
func (e *HTTPRequestError) Error() string {
	return e.message
}
