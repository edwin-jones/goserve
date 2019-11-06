package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/google/uuid"
)

// Server a simple http server that serves get requests from the current working directory
type Server struct {
	port            string
	responseBuilder ResponseBuilder
}

type ResponseBuilder interface {
	Build(rawRequest []byte) ([]byte, error)
}

// New Server constructor
func New(responseBuilder ResponseBuilder) *Server {
	return &Server{
		responseBuilder: responseBuilder,
	}
}

// Serve start serving on the given port
func (s *Server) Serve(port int) string {
	// Listen for incoming connections.
	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))

	if err != nil {
		log.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// Close the listener when the application closes.
	defer l.Close()

	log.Println(fmt.Sprintf("Listening on: %d", port))
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting incoming connection: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go s.handleRequest(conn)
	}
}

// Handles incoming requests.
func (s *Server) handleRequest(conn net.Conn) {

	connectionID := uuid.New()
	log.Println(fmt.Sprintf("Connection %s opened", connectionID))

	// Close the connection last
	defer log.Println(fmt.Sprintf("Connection %s closed", connectionID))
	defer conn.Close()

	// Make a buffer to hold incoming data.
	rawRequest := make([]byte, 1024)

	// Read the incoming connection into the buffer.
	if _, err := conn.Read(rawRequest); err != nil {
		log.Println("Error reading request stream:", err.Error())
		return
	}

	requestString := string(rawRequest)
	requestedUri := strings.Split(requestString, "\n")[0]
	log.Println(fmt.Sprintf("Connection %s is requesting: %s", connectionID, requestedUri))

	responseData, err := s.responseBuilder.Build(rawRequest)

	if err != nil {
		log.Println("Error building response:", err.Error())
		return
	}

	conn.Write(responseData)
}
