// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
	"net/http/httptest"
	"testing"
)

// Tests setting and retrieving a status code
func TestStatusCode(t *testing.T) {
	recorder := httptest.NewRecorder()

	response := NewResponse(recorder)

	if response.GetStatusCode() != 200 {
		t.Fail()
	}

	response.SetStatusCode(404)

	if response.GetStatusCode() != 404 {
		t.Fail()
	}
}
