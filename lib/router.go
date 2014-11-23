// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
    "net/http"
)

type router struct {
    routes []route
}

type route struct {
    name    string
    pattern string
    handler handler
}

type handler func(w http.ResponseWriter, r *http.Request)

// Creates a new router
func NewRouter() router {
    return router{routes: make([]route, 0)}
}

// Adds a new route to router
func (r *router) Add(name string, pattern string, handler handler) {
    route := route{name: name, pattern: pattern, handler: handler}
    r.routes = append(r.routes, route)

    http.HandleFunc(pattern, handler)
}

// Returns all routes available in router
func (r *router) GetRoutes() []route {
    return r.routes
}

// Returns a route from given name
func (r *router) GetRoute(name string) route {
    var result route

    for _, route := range r.routes {
        if (route.name == name) {
            result = route
        }
    }

    return result
}
