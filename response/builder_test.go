package response

import (
	"testing"
	"strings"
	"github.com/edwin-jones/goserve/request"
)

type testFileReader struct{}

func (testFileReader) Read(path string) ([]byte, error) {
	return []byte("some test text"), nil
}

type testRequestParser struct{}

func (testRequestParser) Parse(rawRequest []byte) (string, *request.Error) {
	return "data", nil
}

func TestResponseCanBeBuilt(t *testing.T) {

	fileReader := &testFileReader{}
	requestParser := &testRequestParser{}
	testRequestBytes := []byte("GET test.txt HTTP/1.1")

	Builder := NewBuilder(fileReader, requestParser)

	responseData, err := Builder.Build(testRequestBytes)

	if err != nil {
		t.Errorf("Expected response to be built successfully but got error %s", err.Error())
	}

	stringResponse := string(responseData)

	if !strings.HasPrefix(stringResponse, "HTTP/1.1 200") {
		t.Errorf("Expected response to have prefix 'HTTP/1.1 200' but got %s", stringResponse)
	}
}
