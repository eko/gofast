// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
    "net/http"
)

type response struct {
    http.ResponseWriter
}

// Creates a new response component instance
func NewResponse(res http.ResponseWriter) response {
    return response{res}
}