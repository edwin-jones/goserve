package validation

// httpRequestError A representation of an HTTP error with error response data
type httpRequestError struct {
	message  string
	Response string
}

// newHTTPRequestError Constructor for httpRequestError
func newHTTPRequestError(response string) *httpRequestError {
	return &httpRequestError{
		message:  "An invalid http request has been made.",
		Response: response,
	}
}

func (e *httpRequestError) Error() string {
	return e.message
}
