// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
    "regexp"
)

type router struct {
    routes []route
}

type route struct {
    method     string
    name       string
    pattern    *regexp.Regexp
    handler    handler
}

type RouteLen []route

type handler func()

// Creates a new router component instance
func NewRouter() router {
    return router{routes: make([]route, 0)}
}

// Adds different HTTP methods route
func (r *router) Get(name string, pattern string, handler handler) {
    r.Add("GET", name, pattern, handler)
}

func (r *router) Patch(name string, pattern string, handler handler) {
    r.Add("PATCH", name, pattern, handler)
}

func (r *router) Post(name string, pattern string, handler handler) {
    r.Add("POST", name, pattern, handler)
}

func (r *router) Put(name string, pattern string, handler handler) {
    r.Add("PUT", name, pattern, handler)
}

func (r *router) Delete(name string, pattern string, handler handler) {
    r.Add("DELETE", name, pattern, handler)
}

func (r *router) Options(name string, pattern string, handler handler) {
    r.Add("OPTIONS", name, pattern, handler)
}

func (r *router) Head(name string, pattern string, handler handler) {
    r.Add("HEAD", name, pattern, handler)
}

func (r *router) All(name string, pattern string, handler handler) {
    r.Add("*", name, pattern, handler)
}

// Adds a new route to router
func (r *router) Add(method string, name string, pattern string, handler handler) {
    route := route{method, name, regexp.MustCompile(pattern), handler}
    r.routes = append(r.routes, route)
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

// Sets route fallback (for 404 error pages)
func (r *router) SetFallback(handler handler) {
    r.Add("*", "fallback", "/", handler)
}

// Returns fallback route (for 404 error pages)
func (r *router) GetFallback() route {
    return r.GetRoute("fallback")
}

// Returns a route pattern
func (r *route) GetPattern() *regexp.Regexp {
    return r.pattern
}

// Route sort functions
func (this RouteLen) Len() int {
    return len(this)
}

func (this RouteLen) Less(i, j int) bool {
    return len(this[i].pattern.String()) > len(this[j].pattern.String())
}

func (this RouteLen) Swap(i, j int) {
    this[i], this[j] = this[j], this[i]
}