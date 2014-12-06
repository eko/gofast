// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
    "net/http"
)

type request struct {
    httpRequest *http.Request
    route       *route
    parameters  []parameter
}

type parameter struct {
    name  string
    value string
}

// Creates a new request component instance
func NewRequest(req *http.Request, route route) request {
    return request{req, &route, make([]parameter, 0)}
}

// Returs HTTP request
func (r *request) GetHttpRequest() *http.Request {
    return r.httpRequest
}

// Returns current route
func (r *request) GetRoute() *route {
    return r.route
}

// Adds a request parameter
func (r *request) AddParameter(name string, value string) {
    r.parameters = append(r.parameters, parameter{name, value})
}

// Returns a request parameter from given name
func (r *request) GetParameter(name string) string {
    var result string

    for _, parameter := range r.parameters {
        if (parameter.name == name) {
            result = parameter.value
        }
    }

    return result
}