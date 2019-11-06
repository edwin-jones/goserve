package request

import (
	"strings"

	"github.com/edwin-jones/goserve/status"
)

var supportedFileTypes = []string{".html", ".htm", ".jpeg", ".jpg", ".png", ".gif", ".js", ".css", ".txt"}

// FileChecker interface to fine "file exists" operations
type FileChecker interface {
	Exists(path string) bool
}

// Parser handles parsing http requests for the relevant path.
type Parser struct {
	fileChecker FileChecker
}

// NewParser ctor for parser
func NewParser(fileChecker FileChecker) *Parser {
	return &Parser{
		fileChecker: fileChecker,
	}
}

// Parse this function parses an http request to get the request path.
// Returns an error if the supplied http request isn't valid.
func (p Parser) Parse(rawRequest []byte) (*Data, *Error) {

	data := &Data{}
	requestData := string(rawRequest)

	if len(rawRequest) == 0 {
		return data, newError(status.BadRequest)
	}

	tokens := strings.Split(requestData, " ")
	data.Verb = tokens[0]

	if data.Verb != "GET" {
		return data, newError(status.InvalidHTTPMethod)
	}

	if len(tokens) < 2 {
		return data, newError(status.BadRequest)
	}

	data.Path = tokens[1][1:]

	if len(data.Path) > 512 {
		return data, newError(status.URITooLong)
	}

	if !isFileTypeSupported(data.Path) {
		return data, newError(status.UnsupportedMediaType)
	}

	if !p.fileChecker.Exists(data.Path) {
		return data, newError(status.NotFound)
	}

	data.FileType = tokens[len(tokens)-1]

	return data, nil
}

func isFileTypeSupported(path string) bool {
	var fileTypeSupported = false
	for _, supportedFileType := range supportedFileTypes {
		if strings.HasSuffix(path, supportedFileType) {
			fileTypeSupported = true
			break
		}
	}

	return fileTypeSupported
}
