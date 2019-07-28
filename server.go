package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/google/uuid"
)

// Serve start serving on the given port
func Serve(port string) string {
	// Listen for incoming connections.
	l, err := net.Listen("tcp", "localhost"+":"+port)
	if err != nil {
		log.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// Close the listener when the application closes.
	defer l.Close()

	log.Println(fmt.Sprintf("Listening on: %s", port))
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting incoming connection: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {

	connectionID := uuid.New()
	log.Println(fmt.Sprintf("Opened connection %s", connectionID))

	// Close the connection last
	defer conn.Close()
	defer log.Println(fmt.Sprintf("Closed connection %s", connectionID))

	// Make a buffer to hold incoming data.
	buffer := make([]byte, 1024)

	// Read the incoming connection into the buffer.
	if _, err := conn.Read(buffer); err != nil {
		log.Println("Error reading request stream:", err.Error())
	}

	request := string(buffer)

	if err := ValidateRequest(request); err != nil {
		log.Println(err)
		conn.Write([]byte(err.response))
		return
	}

	log.Println("A successful http request has been handled.")
	conn.Write(BuildSuccessResponse(&request))
}
