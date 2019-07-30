package main

import (
	"testing"
)

func TestBlankRequestValidation(t *testing.T) {
	if err := ValidateRequest(""); err == nil {
		t.Error("expected error response")
	} else {
		if err.response != badRequest {
			t.Error("expected 400 bad request error response")
		}
	}
}

func TestIncompleteRequestValidation(t *testing.T) {
	if err := ValidateRequest("GET foo"); err == nil {
		t.Error("expected error response")
	} else {
		if err.response != badRequest {
			t.Error("expected 400 bad request error response")
		}
	}
}

func TestInvalidMethodValidation(t *testing.T) {
	if err := ValidateRequest("DELETE"); err == nil {
		t.Error("expected error response")
	} else {
		if err.response != invalidHTTPMethod {
			t.Error("expected 405 invalid method error response")
		}
	}
}

func TestUnsupportedFiletypeValidation(t *testing.T) {
	if err := ValidateRequest("GET http://localhost:8080/test.foo"); err == nil {
		t.Error("expected error response")
	} else {
		if err.response != unsupportedMediaType {
			t.Error("expected 415 unsupported filetype error response")
		}
	}
}
