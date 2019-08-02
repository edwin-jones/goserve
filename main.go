package main

import "github.com/edwin-jones/goserve/server"

func main() {
	server := server.New("8080")
	server.Serve()
}
