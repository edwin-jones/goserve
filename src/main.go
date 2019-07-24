package main

import (
	"fmt"
	"net"
	"os"
	"responses"
	"strings"
)

const (
	host     = "localhost"
	port     = "8080"
	protocol = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(protocol, host+":"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// Close the listener when the application closes.
	defer l.Close()

	fmt.Println(fmt.Sprintf("Listening on: %s:%s", host, port))
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	request := string(buf)
	tokens := strings.Split(request, " ")
	verb := tokens[0]

	// Send a response back to person contacting us.
	if verb == "GET" {
		conn.Write([]byte(responses.Success))
	} else {
		conn.Write([]byte(responses.InvalidHttpMethod))
	}

	conn.Close()
}
