// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type Context struct {
	logger   *logrus.Logger
	request  *Request
	response *Response
	route    *Route
}

// Creates a new context component instance
func NewContext() Context {
	return Context{}
}

// Sets Logrus logger instance
func (c *Context) SetLogger(logger *logrus.Logger) {
	c.logger = logger
}

// GetLogger returns a Logrus logger instance
func (c *Context) GetLogger() *logrus.Logger {
	return c.logger
}

// Sets a HTTP request instance
func (c *Context) SetRequest(req *http.Request) {
	request := NewRequest(req)
	c.request = &request
}

// GetRequest returns a HTTP request component instance
func (c *Context) GetRequest() *Request {
	return c.request
}

// Sets a route instance
func (c *Context) SetRoute(route *Route) {
	c.route = route
}

// GetRoute returns a route instance
func (c *Context) GetRoute() *Route {
	return c.route
}

// Sets a HTTP response instance
func (c *Context) SetResponse(res http.ResponseWriter) {
	response := NewResponse(res)
	c.response = &response

	c.AddDefaultHeaders()
}

// GetResponse returns a HTTP response component instance
func (c *Context) GetResponse() *Response {
	return c.response
}

// AddDefaultHeaders adds some defaults headers to send with the response
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
