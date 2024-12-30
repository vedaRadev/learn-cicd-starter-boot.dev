package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	var headers http.Header
	var result string
	var err error

	headers = http.Header{}
	headers.Add("Authorization", "ApiKey test")
	result, err = GetAPIKey(headers)
	if err != nil {
		t.Errorf("errored but shouldn't have: %v\n", err.Error())
		return
	}
	if result != "test" {
		t.Errorf(`incorrect api key: expected "test" but got "%v"\n`, result)
		return
	}

	headers = http.Header{}
	result, err = GetAPIKey(headers)
	if err == nil {
		t.Error("did not error but should have")
		return
	}
	if err != ErrNoAuthHeaderIncluded {
		t.Error("errored but did not return the correct error type, expected ErrNoAuthHeaderIncluded")
		return
	}

	headers = http.Header{}
	headers.Add("Authorization", "Bearer test")
	result, err = GetAPIKey(headers)
	if err == nil {
		t.Error("did not error but should have")
		return
	}
	if !strings.Contains(err.Error(), "malformed") {
		t.Errorf(`errored but got unexpected message: "%v" should have contained "malformed"\n`, err.Error())
		return
	}
}
