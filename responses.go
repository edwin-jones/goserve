package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	successHTMLTemplate = "HTTP/1.1 200 OK\nContent-Type: %s\nContent-Length: %d\n\n"
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

// BuildSuccessResponse Builds a successful HTTP response from the given request
func BuildSuccessResponse(request *string) []byte {
	tokens := strings.Split(*request, " ")
	url := tokens[1][1:]

	return getSuccessResponseBytes(url)
}

func getFileBytes(fileName string) []byte {
	fileBytes, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		panic(err)
	}

	return fileBytes
}

func getSuccessResponseBytes(url string) []byte {
	fileBytes := getFileBytes(url)
	tokens := strings.Split(url, ".")
	fileType := tokens[len(tokens)-1]
	mimeType := mimeTypeMap[fileType]
	response := fmt.Sprintf(successHTMLTemplate, mimeType, len(fileBytes))
	responseBytes := append([]byte(response), fileBytes...)

	return responseBytes
}
