package request

import (
	"strings"
)

var supportedFileTypes = []string{".html", ".htm", ".jpeg", ".jpg", ".png", ".gif", ".js", ".css"}

type FileChecker interface {
	Exists(path string) bool
}

// Parser handles parsing http requests for the relevant path.
type Parser struct {
	fileChecker FileChecker
}

// NewParse ctor for parser
func NewParser(fileChecker FileChecker) *Parser {
	return &Parser{
		fileChecker: fileChecker,
	}
}

// Parse this function parses an http request to get the request path.
// Returns an error if the supplied http request isn't valid.
func (p Parser) Parse(rawRequest string) (string, *Error) {

	path := ""

	if len(rawRequest) == 0 {
		return path, newError(BadRequest)
	}

	tokens := strings.Split(rawRequest, " ")
	verb := tokens[0]

	if verb != "GET" {
		return path, newError(InvalidHTTPMethod)
	}

	if len(tokens) < 2 {
		return path, newError(BadRequest)
	}

	path = tokens[1][1:]

	if len(path) > 512 {
		return path, newError(URITooLong)
	}

	if !isFileTypeSupported(path) {
		return path, newError(UnsupportedMediaType)
	}

	if !p.fileChecker.Exists(path) {
		return path, newError(NotFound)
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
