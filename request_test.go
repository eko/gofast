// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
	"net/http"
	"regexp"
	"testing"
)

// Tests setting and retrieving current route
func TestRoute(t *testing.T) {
	httpRequest := new(http.Request)
	route := Route{"GET", "test", regexp.MustCompile("/test"), func(c Context) {}}

	request := NewRequest(httpRequest, route)

	if request.GetRoute().name != "test" {
		t.Fail()
	}
}

// Tests setting and retrieving request parameters
func TestParameters(t *testing.T) {
	httpRequest := new(http.Request)
	route := Route{"GET", "test", regexp.MustCompile("/test"), func(c Context) {}}

	request := NewRequest(httpRequest, route)

	request.AddParameter("test1", "value1")
	request.AddParameter("test2", "value2")

	if request.GetParameter("test1") != "value1" {
		t.Fail()
	}

	if request.GetParameter("test2") != "value2" {
		t.Fail()
	}
}
