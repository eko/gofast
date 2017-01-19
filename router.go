// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
	"regexp"
)

type Router struct {
	routes []Route
}

type Route struct {
	method  string
	name    string
	pattern *regexp.Regexp
	handler Handler
}

type RouteLen []Route

type Handler func(context Context)

// Creates a new router component instance
func NewRouter() Router {
	return Router{routes: make([]Route, 0)}
}

// Adds different HTTP methods route
func (r *Router) Get(name string, pattern string, handler Handler) {
	r.Add("GET", name, pattern, handler)
}

func (r *Router) Patch(name string, pattern string, handler Handler) {
	r.Add("PATCH", name, pattern, handler)
}

func (r *Router) Post(name string, pattern string, handler Handler) {
	r.Add("POST", name, pattern, handler)
}

func (r *Router) Put(name string, pattern string, handler Handler) {
	r.Add("PUT", name, pattern, handler)
}

func (r *Router) Delete(name string, pattern string, handler Handler) {
	r.Add("DELETE", name, pattern, handler)
}

func (r *Router) Options(name string, pattern string, handler Handler) {
	r.Add("OPTIONS", name, pattern, handler)
}

func (r *Router) Head(name string, pattern string, handler Handler) {
	r.Add("HEAD", name, pattern, handler)
}

func (r *Router) All(name string, pattern string, handler Handler) {
	r.Add("*", name, pattern, handler)
}

// Adds a new route to router
func (r *Router) Add(method string, name string, pattern string, handler Handler) {
	route := Route{method, name, regexp.MustCompile(pattern), handler}
	r.routes = append(r.routes, route)
}

// Returns all routes available in router
func (r *Router) GetRoutes() []Route {
	return r.routes
}

// Returns a Route from given name
func (r *Router) GetRoute(name string) Route {
	var result Route

	for _, route := range r.routes {
		if route.name == name {
			result = route
		}
	}

	return result
}

// Sets route fallback (for 404 error pages)
func (r *Router) SetFallback(handler Handler) {
	r.Add("*", "fallback", "/", handler)
}

// Returns fallback route (for 404 error pages)
func (r *Router) GetFallback() Route {
	return r.GetRoute("fallback")
}

// Returns a route pattern
func (r *Route) GetPattern() *regexp.Regexp {
	return r.pattern
}

// Sets a route handler
func (r *Route) SetHandler(handler Handler) {
	r.handler = handler
}

// Returns a route handler
func (r *Route) GetHandler() Handler {
	return r.handler
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
