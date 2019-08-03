package response

import (
	"bytes"
	"fmt"
	"github.com/edwin-jones/goserve/request"
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

type FileReader interface {
	Read(path string) ([]byte, error)
}

// Response constructs byte responses to http requests
type Response struct {
	fileReader FileReader
}

func New(fileReader FileReader) *Response {
	return &Response{fileReader: fileReader}
}

// Build an http response based on the status code
func (r Response) Build(StatusCode request.StatusCode, path string) ([]byte, error) {

	var err error
	var responseBytes []byte

	switch StatusCode {
	case request.Success:
		responseBytes, err = r.buildSuccess(path)
	case request.BadRequest:
		responseBytes = []byte(badRequestResponse)
	case request.NotFound:
		responseBytes = []byte(notFoundResponse)
	case request.URITooLong:
		responseBytes = []byte(uriTooLongResponse)
	case request.UnsupportedMediaType:
		responseBytes = []byte(unsupportedMediaTypeResponse)
	case request.InvalidHTTPMethod:
		responseBytes = []byte(invalidHTTPMethodResponse)
	default:
		responseBytes = []byte(badRequestResponse)
	}

	return responseBytes, err
}

// BuildSuccess Builds a successful HTTP response from an http request path
func (r Response) buildSuccess(path string) ([]byte, error) {

	var err error
	var fileBytes []byte
	var buffer bytes.Buffer

	tokens := strings.Split(path, ".")
	fileType := tokens[len(tokens)-1]
	mimeType := mimeTypeMap[fileType]
	fileBytes, err = r.fileReader.Read(path)

	_, err = fmt.Fprintf(&buffer, successHTMLTemplate, mimeType, len(fileBytes))
	_, err = buffer.Write(fileBytes)

	return buffer.Bytes(), err
}
