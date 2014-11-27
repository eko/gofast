// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import(
    "net/http"
    "os"
    "log"
)

const (
    VERSION string = "1.0-beta"
)

type gofast struct {
    logger     *log.Logger
    router     router
    templating templating
}

type requestHandler func()

// Bootstraps a new instance
func Bootstrap() gofast {
    log.Printf("gofast v%s", VERSION)

    logger     := log.New(os.Stdout, "[gofast]", 0)
    templating := NewTemplating()
    router     := NewRouter(templating)

    return gofast{logger, router, templating}
}

// Handles HTTP requests
func (g *gofast) Handle() {
    http.ListenAndServe(":8080", nil)
}

// Returns router component
func (g *gofast) GetRouter() router {
    return g.router
}

// Returns templating component
func (g *gofast) GetTemplating() templating {
    return g.templating
}
