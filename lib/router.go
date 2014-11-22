// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

type router struct {
    routes []route
}

type route struct {
    name    string
    pattern string
}

// Creates a new router
func NewRouter() router {
	  return router{routes: make([]route, 0)}
}

// Adds a new route to router
func (r *router) AddRoute(name string, pattern string) {
    route := route{name: name, pattern: pattern}
    r.routes = append(r.routes, route)
}

// Returns all routes available in router
func (r *router) GetRoutes() []route {
    return r.routes
}