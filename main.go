package main

import (
	"github.com/edwin-jones/goserve/file"
	"github.com/edwin-jones/goserve/request"
	"github.com/edwin-jones/goserve/response"
	"github.com/edwin-jones/goserve/server"
)

func main() {
	fileHandler := file.Handler{}
	requestParser := request.NewParser(fileHandler)
	responseBuilder := response.NewBuilder(fileHandler)

	server := server.New(requestParser, responseBuilder)
	server.Serve(8080)
}
