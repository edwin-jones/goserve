package parser

import (
	"testing"
)

func TestBlankRequestValidation(t *testing.T) {
	if _, err := Parse(""); err == nil {
		t.Error("expected error response")
	} else {
		if err.Response != badRequest {
			t.Error("expected 400 bad request error response")
		}
	}
}

func TestIncompleteRequestValidation(t *testing.T) {
	if _, err := Parse("GET foo"); err == nil {
		t.Error("expected error response")
	} else {
		if err.Response != badRequest {
			t.Error("expected 400 bad request error response")
		}
	}
}

func TestInvalidMethodValidation(t *testing.T) {
	if _, err := Parse("DELETE"); err == nil {
		t.Error("expected error response")
	} else {
		if err.Response != invalidHTTPMethod {
			t.Error("expected 405 invalid method error response")
		}
	}
}

func TestUnsupportedFiletypeValidation(t *testing.T) {
	if _, err := Parse("GET http://localhost:8080/test.foo"); err == nil {
		t.Error("expected error response")
	} else {
		if err.Response != unsupportedMediaType {
			t.Error("expected 415 unsupported filetype error response")
		}
	}
}
