package request

import (
	"testing"

	"github.com/dchest/uniuri"
	"github.com/edwin-jones/goserve/status"
)

type testFileChecker struct {
	exists bool
}

func (t testFileChecker) Exists(path string) bool {
	return t.exists
}

func TestStatusValidation(t *testing.T) {
	testCases := []struct {
		request            string
		expectedStatusCode status.Code
		fileExists         bool
	}{
		{
			request:            "",
			expectedStatusCode: status.BadRequest,
			fileExists:         true,
		},
		{
			request:            "GET foo",
			expectedStatusCode: status.BadRequest,
			fileExists:         true,
		},
		{
			request:            "GET test.foo",
			expectedStatusCode: status.UnsupportedMediaType,
			fileExists:         true,
		},
		{
			request:            "GET test.html" + uniuri.NewLen(1000),
			expectedStatusCode: status.URITooLong,
			fileExists:         true,
		},
		{
			request:            "DELETE",
			expectedStatusCode: status.InvalidHTTPMethod,
			fileExists:         true,
		},
		{
			request:            "GET test.html",
			expectedStatusCode: status.Success,
			fileExists:         true,
		},
		{
			request:            "GET test.html",
			expectedStatusCode: status.NotFound,
			fileExists:         false,
		},
	}

	for _, c := range testCases {
		fileChecker := testFileChecker{exists: c.fileExists}
		parser := NewParser(fileChecker)

		_, code := parser.Parse([]byte(c.request))
		if code != c.expectedStatusCode {
			t.Errorf("Expected status code %d, got %d", c.expectedStatusCode, code)
		}
	}
}
