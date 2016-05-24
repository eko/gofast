// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
	"net/http"
)

type Context struct {
	request  *Request
	response *Response
}

// Creates a new context component instance
func NewContext() Context {
	return Context{}
}

// Sets a HTTP request instance
func (c *Context) SetRequest(req *http.Request, route Route) {
	request := NewRequest(req, route)
	c.request = &request
}

// Returns a HTTP request component instance
func (c *Context) GetRequest() *Request {
	return c.request
}

// Sets a HTTP response instance
func (c *Context) SetResponse(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "text/html; charset: utf-8")

	response := NewResponse(res)
	c.response = &response
}

// Returns a HTTP response component instance
func (c *Context) GetResponse() *Response {
	return c.response
}
