package main

import "github.com/edwin-jones/goserve/server"

const (
	port = "8080"
)

func main() {
	server.Serve(port)
}
