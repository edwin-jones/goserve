package response

import (
	"fmt"
	"github.com/edwin-jones/goserve/request"
	"io"
	"io/ioutil"
	"strings"
)

const (
	successHTMLTemplate          = "HTTP/1.1 200 OK\nContent-Type: %s\nContent-Length: %d\n\n"
	badRequestResponse           = "HTTP/1.1 400 Bad Request\nContent-Type: text/plain\nContent-Length: 15\n\n400 Bad Request"
	notFoundResponse             = "HTTP/1.1 404 Not Found\nContent-Type: text/plain\nContent-Length: 13\n\n404 Not Found"
	invalidHTTPMethodResponse    = "HTTP/1.1 405 Method Not Allowed\nAllow: GET\nContent-Type: text/plain\nContent-Length: 22\n\n405 Method Not Allowed"
	uriTooLongResponse           = "HTTP/1.1 414 URI Too Long\nContent-Type: text/plain\nContent-Length: 16\n\n414 URI Too Long"
	unsupportedMediaTypeResponse = "HTTP/1.1 415 Unsupported Media Type\nContent-Type: text/plain\nContent-Length: 26\n\n415 Unsupported Media Type"
)

var mimeTypeMap = map[string]string{
	"html": "text/html",
	"htm":  "text/html",
	"css":  "text/css",
	"jpeg": "image/jpeg",
	"jpg":  "image/jpeg",
	"gif":  "image/gif",
	"png":  "image/png",
	"js":   "application/javascript",
}

// Response constructs byte responses to http requests
type Response struct {
	writer io.Writer
	reader io.Reader
}

// New constructor for response
func New(writer io.Writer, reader io.Reader) *Response {
	return &Response{writer: writer, reader: reader}
}

// BuildSuccess Builds a successful HTTP response from an http request path
func (r *Response) BuildSuccess(path string) {

	var err error
	var fileBytes []byte

	fileBytes, err = ioutil.ReadAll(r.reader)
	tokens := strings.Split(path, ".")
	fileType := tokens[len(tokens)-1]
	mimeType := mimeTypeMap[fileType]

	_, err = fmt.Fprintf(r.writer, successHTMLTemplate, mimeType, len(fileBytes))
	_, err = r.writer.Write(fileBytes)

	if err != nil {
		panic(err)
	}
}

// BuildError Builds an error HTTP response from the an http error code
func (r Response) BuildError(errorCode request.ErrorCode) {

	var err error

	switch errorCode {
	case request.BadRequest:
		_, err = r.writer.Write([]byte(badRequestResponse))
	case request.NotFound:
		_, err = r.writer.Write([]byte(notFoundResponse))
	case request.URITooLong:
		_, err = r.writer.Write([]byte(uriTooLongResponse))
	case request.UnsupportedMediaType:
		_, err = r.writer.Write([]byte(unsupportedMediaTypeResponse))
	case request.InvalidHTTPMethod:
		_, err = r.writer.Write([]byte(invalidHTTPMethodResponse))
	default:
		_, err = r.writer.Write([]byte(badRequestResponse))
	}

	if err != nil {
		panic(err)
	}
}
