// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
    "net/http"
    "time"
    "log"
)

type router struct {
    templating templating
    routes     []route
}

type route struct {
    method     string
    name       string
    pattern    string
    handler    handler
}

type handler func(w http.ResponseWriter, r *http.Request)

// Creates a new router component instance
func NewRouter(t templating) router {
    return router{templating: t, routes: make([]route, 0)}
}

// Add different HTTP methods route
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
    route := route{method, name, pattern, handler}
    r.routes = append(r.routes, route)

    requestHandler := func(w http.ResponseWriter, r *http.Request) {
        if (r.Method == method) {
            t1 := time.Now()
            handler(w, r)
            t2 := time.Now()

            log.Printf("[%s] %q (time: %v)\n", r.Method, r.URL.String(), t2.Sub(t1))
        }
    }

    http.HandleFunc(pattern, requestHandler)
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
