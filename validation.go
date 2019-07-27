package main

import (
	"os"
	"strings"
)

const (
	notFound             = "HTTP/1.1 404 Not Found\nContent-Type: text/plain\nContent-Length: 13\n\n404 Not Found"
	invalidHTTPMethod    = "HTTP/1.1 405 Method Not Allowed\nAllow: GET\nContent-Type: text/plain\nContent-Length: 22\n\n405 Method Not Allowed"
	unsupportedMediaType = "HTTP/1.1 415 Unsupported Media Type\nContent-Type: text/plain\nContent-Length: 26\n\n415 Unsupported Media Type"
)

var supportedFileTypes = []string{".html", ".htm"}

// ValidateRequest this function returns an error if the supplied http request isn't valid
func ValidateRequest(rawRequest string) *HTTPRequestError {

	tokens := strings.Split(rawRequest, " ")
	verb := tokens[0]

	if verb != "GET" {
		error := NewHTTPRequestError([]byte(invalidHTTPMethod))
		return error
	}

	url := tokens[1][1:]

	if _, fileError := os.Stat(url); os.IsNotExist(fileError) {
		error := NewHTTPRequestError([]byte(notFound))
		return error
	}

	if !isFileTypeSupported(&url) {
		return NewHTTPRequestError([]byte(unsupportedMediaType))
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
