// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
	"net/http"
	"testing"
)

// TestParameters tests setting and retrieving request parameters
func TestParameters(t *testing.T) {
	httpRequest := new(http.Request)
	request := NewRequest(httpRequest)

	request.AddParameter("test1", "value1")
	request.AddParameter("test2", "value2")

	if request.GetParameter("test1") != "value1" {
		t.Fail()
	}

	if request.GetParameter("test2") != "value2" {
		t.Fail()
	}
}

// TestGetHeader tests retrieving a header
func TestGetHeader(t *testing.T) {
	httpRequest, _ := http.NewRequest("GET", "/", nil)
	httpRequest.Header.Set("X-Test-Header", "yes")

	request := NewRequest(httpRequest)

	if request.GetHeader("X-Test-Header") != "yes" {
		t.Fail()
	}
}
