// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
    "net/http"
    "testing"
)

// Tests directory setter/getter
func TestSetDirectory(t *testing.T) {
    c := NewContext()
    templating := c.GetTemplating()

    templating.SetDirectory("../views")

    if ("../views" != templating.GetDirectory()) {
        t.Fail()
    }
}

// Tests rendering a view via pongo2 library
func TestRender(t *testing.T) {
    c := Bootstrap().GetContext()

    templating := c.GetTemplating()
    templating.SetDirectory("../views")

    http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
        templating.Render(c, "index.html")
    })
}