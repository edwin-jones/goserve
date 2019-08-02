package response

import (
	"fmt"
	"github.com/edwin-jones/goserve/request"
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

// Builder constructs byte responses to http requests
type Builder struct {
}

// BuildSuccess Builds a successful HTTP response from an http request path
func (Builder) BuildSuccess(path string) []byte {

	fileBytes := getFileBytes(path)
	tokens := strings.Split(path, ".")
	fileType := tokens[len(tokens)-1]
	mimeType := mimeTypeMap[fileType]
	response := fmt.Sprintf(successHTMLTemplate, mimeType, len(fileBytes))
	responseBytes := append([]byte(response), fileBytes...)

	return responseBytes
}

// BuildError Builds an error HTTP response from the an http error code
func (Builder) BuildError(errorCode request.ErrorCode) []byte {

	switch errorCode {
	case request.BadRequest:
		return []byte(badRequestResponse)
	case request.NotFound:
		return []byte(notFoundResponse)
	case request.URITooLong:
		return []byte(uriTooLongResponse)
	case request.UnsupportedMediaType:
		return []byte(unsupportedMediaTypeResponse)
	case request.InvalidHTTPMethod:
		return []byte(invalidHTTPMethodResponse)
	default:
		return []byte(badRequestResponse)
	}
}

func getFileBytes(fileName string) []byte {
	fileBytes, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		panic(err)
	}

	return fileBytes
}
