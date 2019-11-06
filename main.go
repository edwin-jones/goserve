package main

import (
	"flag"

	"github.com/edwin-jones/goserve/file"
	"github.com/edwin-jones/goserve/request"
	"github.com/edwin-jones/goserve/response"
	"github.com/edwin-jones/goserve/server"
)

func main() {

	var port int
	flag.IntVar(&port, "p", 8080, "the port to run the http server on")
	flag.Parse()

	fileHandler := file.Handler{}
	requestParser := request.NewParser(fileHandler)
	responseBuilder := response.NewBuilder(fileHandler)

	server := server.New(requestParser, responseBuilder)
	server.Serve(port)
}
