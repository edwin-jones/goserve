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
	if verb != "GET" {
		conn.Write([]byte(responses.InvalidHttpMethod))
	} else {

		url := tokens[1][1:]
		files := getDirectoryFileNames()

		if stringInSlice(url, files) {
			fileBytes := getFileBytes(url)
			response := "HTTP/1.1 200 OK\nContent-Type: text/html\nContent-Length: " + fmt.Sprint(len(fileBytes)) + "\n\n"
			responseBytes := append([]byte(response), fileBytes...)
			conn.Write(responseBytes)
		} else {
			conn.Write([]byte(responses.NotFound))
		}
	}

	conn.Close()
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

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getFileBytes(fileName string) []byte {
	fileBytes, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		panic(err)
	}

	return fileBytes
}
