// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
	"net/http"
)

type Response struct {
	http.ResponseWriter
	statusCode int
}

// NewResponse creates a new Response component instance
func NewResponse(res http.ResponseWriter) Response {
	return Response{res, 200}
}

// SetStatusCode sets Response status code
func (r *Response) SetStatusCode(statusCode int) {
	r.WriteHeader(statusCode)
	r.statusCode = statusCode
}

// GetStatusCode returns Response status code
func (r *Response) GetStatusCode() int {
	return r.statusCode
}
