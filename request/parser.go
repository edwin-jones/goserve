package request

import (
	"os"
	"strings"
)

var supportedFileTypes = []string{".html", ".htm", ".jpeg", ".jpg", ".png", ".gif", ".js", ".css"}

// Parser handles parsing http requests for the relevant path.
type Parser struct {
}

// Parse this function parses an http request to get the request path.
// Returns an error if the supplied http request isn't valid.
func (Parser) Parse(rawRequest string) (string, *Error) {

	url := ""

	if len(rawRequest) == 0 {
		return url, newError(BadRequest)
	}

	tokens := strings.Split(rawRequest, " ")
	verb := tokens[0]

	if verb != "GET" {
		return url, newError(InvalidHTTPMethod)
	}

	if len(tokens) < 2 {
		return url, newError(BadRequest)
	}

	url = tokens[1][1:]

	if len(url) > 512 {
		return url, newError(URITooLong)
	}

	if !isFileTypeSupported(&url) {
		return url, newError(UnsupportedMediaType)
	}

	if _, fileError := os.Stat(url); os.IsNotExist(fileError) {
		return url, newError(NotFound)
	}

	return url, nil
}

func isFileTypeSupported(url *string) bool {
	var fileTypeSupported = false
	for _, supportedFileType := range supportedFileTypes {
		if strings.HasSuffix(*url, supportedFileType) {
			fileTypeSupported = true
			break
		}
	}

	return fileTypeSupported
}
