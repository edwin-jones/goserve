package main

import (
	"fmt"
	"io/ioutil"
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

	fmt.Println("Opening connection")

	// Close the connection last
	defer conn.Close()
	defer fmt.Println("Closing connection")

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

	if verb != "GET" {
		conn.Write([]byte(responses.InvalidHttpMethod))
		return
	}

	url := tokens[1][1:]

	if _, fileError := os.Stat(url); os.IsNotExist(fileError) {
		conn.Write([]byte(responses.NotFound))
		return
	}

	if strings.HasSuffix(url, "html") || strings.HasSuffix(url, "htm") {
		conn.Write(getSuccessResponseBytes(url))
	} else {
		conn.Write([]byte(responses.UnsupportedMediaType))
	}
}

func getDirectoryFileNames() []string {
	files, err := ioutil.ReadDir("./")

	if err != nil {
		panic(err)
	}

	var results []string
	for _, f := range files {
		results = append(results, f.Name())
	}

	return results
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
	response := responses.SuccessHtmlPrefix + fmt.Sprint(len(fileBytes)) + "\n\n"
	responseBytes := append([]byte(response), fileBytes...)

	return responseBytes
}
