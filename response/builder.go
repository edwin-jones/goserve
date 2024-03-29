package response

import (
	"bytes"
	"fmt"

	"github.com/edwin-jones/goserve/request"
	"github.com/edwin-jones/goserve/status"
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

type FileReader interface {
	Read(path string) ([]byte, error)
}

type RequestParser interface {
	Parse(rawRequest []byte) (request.Data, *request.Error)
}

// Builder constructs byte responses to http requests
type Builder struct {
	fileReader FileReader
}

// NewBuilder ctor for Builder
func NewBuilder(fileReader FileReader) *Builder {
	return &Builder{
		fileReader: fileReader,
	}
}

// Build an http response based on the status code
func (b Builder) Build(data *request.Data, statusCode status.Code) ([]byte, error) {

	var responseBytes []byte
	var err error

	switch statusCode {
	case status.Success:
		responseBytes, err = b.buildSuccess(data)
	case status.BadRequest:
		responseBytes = []byte(badRequestResponse)
	case status.NotFound:
		responseBytes = []byte(notFoundResponse)
	case status.URITooLong:
		responseBytes = []byte(uriTooLongResponse)
	case status.UnsupportedMediaType:
		responseBytes = []byte(unsupportedMediaTypeResponse)
	case status.InvalidHTTPMethod:
		responseBytes = []byte(invalidHTTPMethodResponse)
	default:
		responseBytes = []byte(badRequestResponse)
	}

	return responseBytes, err
}

// BuildSuccess Builds a successful HTTP response from an http request path
func (b Builder) buildSuccess(data *request.Data) ([]byte, error) {

	var err error
	var fileBytes []byte
	var buffer bytes.Buffer

	mimeType := mimeTypeMap[data.FileType]
	fileBytes, err = b.fileReader.Read(data.Path)

	_, err = fmt.Fprintf(&buffer, successHTMLTemplate, mimeType, len(fileBytes))
	_, err = buffer.Write(fileBytes)

	return buffer.Bytes(), err
}
