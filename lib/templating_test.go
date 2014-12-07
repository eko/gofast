// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
    "net/http"
    "testing"
)

// Tests directory setter/getter
func TestSetDirectories(t *testing.T) {
    c := NewContext()
    templating := c.GetTemplating()

    templating.SetAssetsDirectory("../assets")
    templating.SetViewsDirectory("../views")

    if ("../assets" != templating.GetAssetsDirectory()) {
        t.Fail()
    }

    if ("../views" != templating.GetViewsDirectory()) {
        t.Fail()
    }
}

// Tests rendering a view via pongo2 library
func TestRender(t *testing.T) {
    c := Bootstrap().GetContext()

    templating := c.GetTemplating()
    templating.SetViewsDirectory("../views")

    http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
        templating.Render(c, "index.html")
    })
}