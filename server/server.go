package server

import (
	"fmt"
	"github.com/edwin-jones/goserve/file"
	"github.com/edwin-jones/goserve/request"
	"github.com/edwin-jones/goserve/response"
	"github.com/google/uuid"
	"log"
	"net"
	"os"
)

// Server a simple http server that serves get requests from the current working directory
type Server struct {
	port string
}

// New Server constructor
func New(port string) *Server {
	return &Server{port: port}
}

// Serve start serving on the given port
func (server *Server) Serve() string {
	// Listen for incoming connections.
	l, err := net.Listen("tcp", "localhost"+":"+server.port)

	if err != nil {
		log.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// Close the listener when the application closes.
	defer l.Close()

	log.Println(fmt.Sprintf("Listening on: %s", server.port))
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
	defer log.Println(fmt.Sprintf("Closed connection %s", connectionID))
	defer conn.Close()

	// Make a buffer to hold incoming data.
	requestBuffer := make([]byte, 1024)

	// Read the incoming connection into the buffer.
	if _, err := conn.Read(requestBuffer); err != nil {
		log.Println("Error reading request stream:", err.Error())
		return
	}

	requestData := string(requestBuffer)

	fileHandler := file.Handler{}
	parser := request.NewParser(fileHandler)
	path, err := parser.Parse(requestData)

	res := response.New(fileHandler)

	statusCode := request.Success
	if err != nil {
		log.Println(err)
		statusCode = err.StatusCode
	}

	responseData, _ := res.Build(statusCode, path)
	conn.Write(responseData)
}
