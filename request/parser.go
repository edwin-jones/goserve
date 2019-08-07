package request

import (
	"strings"

	"github.com/edwin-jones/goserve/status"
)

var supportedFileTypes = []string{".html", ".htm", ".jpeg", ".jpg", ".png", ".gif", ".js", ".css"}

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
func (p Parser) Parse(rawRequest []byte) (string, *Error) {

	path := ""
	requestData := string(rawRequest)

	if len(rawRequest) == 0 {
		return path, newError(status.BadRequest)
	}

	tokens := strings.Split(requestData, " ")
	verb := tokens[0]

	if verb != "GET" {
		return path, newError(status.InvalidHTTPMethod)
	}

	if len(tokens) < 2 {
		return path, newError(status.BadRequest)
	}

	path = tokens[1][1:]

	if len(path) > 512 {
		return path, newError(status.URITooLong)
	}

	if !isFileTypeSupported(path) {
		return path, newError(status.UnsupportedMediaType)
	}

	if !p.fileChecker.Exists(path) {
		return path, newError(status.NotFound)
	}

	return path, nil
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
