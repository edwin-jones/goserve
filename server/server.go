package server

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/edwin-jones/goserve/request"
	"github.com/edwin-jones/goserve/status"
	"github.com/google/uuid"
)

// Server a simple http server that serves get requests from the current working directory
type Server struct {
	port            string
	requestParser   RequestParser
	responseBuilder ResponseBuilder
}

type ResponseBuilder interface {
	Build(data request.Data, statusCode status.Code) ([]byte, error)
}

type RequestParser interface {
	Parse(rawRequest []byte) (request.Data, *request.Error)
}

// New Server constructor
func New(requestParser RequestParser, responseBuilder ResponseBuilder) *Server {
	return &Server{
		requestParser:   requestParser,
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
	if _, streamError := conn.Read(rawRequest); streamError != nil {
		log.Println(fmt.Sprintf("Connection %s %s", connectionID, streamError.Error()))
		return
	}

	data, requestErr := s.requestParser.Parse(rawRequest)

	statusCode := status.Success

	if requestErr != nil {
		statusCode = requestErr.StatusCode
	}

	responseData, responseErr := s.responseBuilder.Build(data, statusCode)

	if responseErr != nil {
		log.Println(fmt.Sprintf("Connection %s %s", connectionID, responseErr.Error()))
		return
	}

	conn.Write(responseData)

	log.Println(fmt.Sprintf("Connection %s %s %d %s", connectionID, data.Verb, statusCode, data.Path))
}
