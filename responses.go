package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	successHTMLPrefix = "HTTP/1.1 200 OK\nContent-Type: text/html\nContent-Length: "
)

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
	response := successHTMLPrefix + fmt.Sprint(len(fileBytes)) + "\n\n"
	responseBytes := append([]byte(response), fileBytes...)

	return responseBytes
}
