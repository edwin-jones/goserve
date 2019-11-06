package response

import (
	"strings"
	"testing"

	"github.com/edwin-jones/goserve/request"
	"github.com/edwin-jones/goserve/status"
)

type testFileReader struct{}

func (testFileReader) Read(path string) ([]byte, error) {
	return []byte("some test text"), nil
}

func TestResponseCanBeBuilt(t *testing.T) {

	fileReader := &testFileReader{}

	Builder := NewBuilder(fileReader)

	responseData, err := Builder.Build(request.Data{}, status.Success)

	if err != nil {
		t.Errorf("Expected response to be built successfully but got error %s", err.Error())
	}

	stringResponse := string(responseData)

	if !strings.HasPrefix(stringResponse, "HTTP/1.1 200") {
		t.Errorf("Expected response to have prefix 'HTTP/1.1 200' but got %s", stringResponse)
	}
}
