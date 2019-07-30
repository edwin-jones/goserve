package main

import (
	"os"
	"strings"
)

const (
	badRequest           = "HTTP/1.1 400 Bad Request\nContent-Type: text/plain\nContent-Length: 15\n\n400 Bad Request"
	notFound             = "HTTP/1.1 404 Not Found\nContent-Type: text/plain\nContent-Length: 13\n\n404 Not Found"
	invalidHTTPMethod    = "HTTP/1.1 405 Method Not Allowed\nAllow: GET\nContent-Type: text/plain\nContent-Length: 22\n\n405 Method Not Allowed"
	unsupportedMediaType = "HTTP/1.1 415 Unsupported Media Type\nContent-Type: text/plain\nContent-Length: 26\n\n415 Unsupported Media Type"
)

var supportedFileTypes = []string{".html", ".htm", ".jpeg", ".jpg", ".png", ".gif", ".js", ".css"}

// ValidateRequest this function returns an error if the supplied http request isn't valid
func ValidateRequest(rawRequest string) *HTTPRequestError {
	if len(rawRequest) == 0 {
		return NewHTTPRequestError(badRequest)
	}

	tokens := strings.Split(rawRequest, " ")
	verb := tokens[0]

	if verb != "GET" {
		return NewHTTPRequestError(invalidHTTPMethod)
	}

	if len(tokens) < 2 {
		return NewHTTPRequestError(badRequest)
	}

	url := tokens[1][1:]

	if !isFileTypeSupported(&url) {
		return NewHTTPRequestError(unsupportedMediaType)
	}

	if _, fileError := os.Stat(url); os.IsNotExist(fileError) {
		return NewHTTPRequestError(notFound)
	}

	return nil
}

func isFileTypeSupported(url *string) bool {
	var fileTypeSupported = false
	for _, supportedFileType := range supportedFileTypes {
		if strings.HasSuffix(*url, supportedFileType) {
			fileTypeSupported = true
			break
		}
	}

	return fileTypeSupported
}
