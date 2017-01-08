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
	response := NewResponse(res)
	c.response = &response

	c.AddDefaultHeaders()
}

// Returns a HTTP response component instance
func (c *Context) GetResponse() *Response {
	return c.response
}

// Adds some defaults headers to send with the response
func (c *Context) AddDefaultHeaders() {
	request := c.GetRequest()
	response := c.GetResponse()

	response.Header().Set("Content-Type", "text/html; charset: utf-8")

	if origin := request.GetHeader("Origin"); origin != "" {
		response.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		response.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		response.Header().Set("Access-Control-Allow-Origin", origin)
	}
}
