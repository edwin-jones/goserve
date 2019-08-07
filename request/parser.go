package request

import (
	"strings"

	"github.com/edwin-jones/goserve/status"
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
func (p Parser) Parse(rawRequest string) (string, status.Code) {

	path := ""

	if len(rawRequest) == 0 {
		return path, status.BadRequest
	}

	tokens := strings.Split(rawRequest, " ")
	verb := tokens[0]

	if verb != "GET" {
		return path, status.InvalidHTTPMethod
	}

	if len(tokens) < 2 {
		return path, status.BadRequest
	}

	path = tokens[1][1:]

	if len(path) > 512 {
		return path, status.URITooLong
	}

	if !isFileTypeSupported(path) {
		return path, status.UnsupportedMediaType
	}

	if !p.fileChecker.Exists(path) {
		return path, status.NotFound
	}

	return path, status.Success
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
