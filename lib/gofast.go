// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import(
    "log"
)

const (
    PORT    string = ":8080"
    VERSION string = "1.0-beta"
)

type gofast struct {
    context *context
}

type requestHandler func()

// Bootstraps a new instance
func Bootstrap() *gofast {
    log.Printf("gofast v%s", VERSION)

    context := NewContext()

    return &gofast{&context}
}

// Returns context instance
func (g *gofast) GetContext() *context {
    return g.context
}