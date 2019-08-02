package response

import (
	"fmt"
	"github.com/edwin-jones/goserve/request"
	"io/ioutil"
	"strings"
)

const (
	successHTMLTemplate  = "HTTP/1.1 200 OK\nContent-Type: %s\nContent-Length: %d\n\n"
	badRequest           = "HTTP/1.1 400 Bad Request\nContent-Type: text/plain\nContent-Length: 15\n\n400 Bad Request"
	notFound             = "HTTP/1.1 404 Not Found\nContent-Type: text/plain\nContent-Length: 13\n\n404 Not Found"
	invalidHTTPMethod    = "HTTP/1.1 405 Method Not Allowed\nAllow: GET\nContent-Type: text/plain\nContent-Length: 22\n\n405 Method Not Allowed"
	uriTooLong           = "HTTP/1.1 414 URI Too Long\nContent-Type: text/plain\nContent-Length: 16\n\n414 URI Too Long"
	unsupportedMediaType = "HTTP/1.1 415 Unsupported Media Type\nContent-Type: text/plain\nContent-Length: 26\n\n415 Unsupported Media Type"
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

type Builder interface {
	Build(string) []byte
}

type ResponseBuilder struct {
}

// BuildSuccess Builds a successful HTTP response from the given request
func (ResponseBuilder) BuildSuccess(path string) []byte {

	fileBytes := getFileBytes(path)
	tokens := strings.Split(path, ".")
	fileType := tokens[len(tokens)-1]
	mimeType := mimeTypeMap[fileType]
	response := fmt.Sprintf(successHTMLTemplate, mimeType, len(fileBytes))
	responseBytes := append([]byte(response), fileBytes...)

	return responseBytes
}

// BuildError Builds an error HTTP response from the given error code
func (ResponseBuilder) BuildError(errorCode request.ErrorCode) []byte {

	switch errorCode {
	case request.BadRequest:
		return []byte(badRequest)
	case request.NotFound:
		return []byte(notFound)
	case request.URITooLong:
		return []byte(uriTooLong)
	case request.UnsupportedMediaType:
		return []byte(unsupportedMediaType)
	case request.InvalidHTTPMethod:
		return []byte(invalidHTTPMethod)
	default:
		return []byte(badRequest)
	}
}

func getFileBytes(fileName string) []byte {
	fileBytes, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		panic(err)
	}

	return fileBytes
}
