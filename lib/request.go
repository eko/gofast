// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
    "net/http"
)

type request struct {
    *http.Request
    route route
}

// Creates a new request component instance
func NewRequest(req *http.Request, route route) request {
    return request{req, route}
}

// Returns current route
func (r *request) GetRoute() route {
    return r.route
}