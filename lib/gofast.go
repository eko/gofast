// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import(
    "net/http"
    "log"
    "os"
    "time"
)

const (
    PORT    string = ":8080"
    VERSION string = "1.0-beta"
)

type gofast struct {
    logger     *log.Logger
    router     *router
    templating templating
}

type requestHandler func()

// Bootstraps a new instance
func Bootstrap() gofast {
    log.Printf("gofast v%s", VERSION)

    logger     := log.New(os.Stdout, "[gofast]", 0)
    templating := NewTemplating()
    router     := NewRouter()

    return gofast{logger, &router, templating}
}

// Returns router component
func (g *gofast) GetRouter() *router {
    return g.router
}

// Returns templating component
func (g *gofast) GetTemplating() templating {
    return g.templating
}

// Handles HTTP requests
func (g *gofast) Handle() {
    for _, route := range g.GetRouter().GetRoutes() {
        http.HandleFunc(route.pattern, g.HandleRequest(route))
    }

    http.ListenAndServe(PORT, nil)
}

// Handles an HTTP request by logging calls
func (g *gofast) HandleRequest(route route) http.HandlerFunc {
    return func (res http.ResponseWriter, req *http.Request) {
        if (req.Method == route.method) {
            startTime := time.Now()
            route.handler(res, req)
            stopTime := time.Now()

            log.Printf("[%s] %q (time: %v)\n", req.Method, req.URL.String(), stopTime.Sub(startTime))
        }
    }
}