// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
    "net/http"
)

type response struct {
    http.ResponseWriter
    statusCode int
}

// Creates a new response component instance
func NewResponse(res http.ResponseWriter) response {
    return response{res, 200}
}

// Sets response status code
func (r *response) SetStatusCode(statusCode int) {
    r.WriteHeader(statusCode)
    r.statusCode = statusCode
}

// Returns response status code
func (r *response) GetStatusCode() int {
    return r.statusCode
}