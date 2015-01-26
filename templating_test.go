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
    templating := NewTemplating()

    templating.SetAssetsDirectory("../")
    templating.SetViewsDirectory("../")

    if ("../" != templating.GetAssetsDirectory()) {
        t.Fail()
    }

    if ("../" != templating.GetViewsDirectory()) {
        t.Fail()
    }
}

// Tests rendering a view via pongo2 library
func TestRender(t *testing.T) {
    templating := NewTemplating()
    templating.SetViewsDirectory("../")

    context := NewContext()

    http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
        templating.Render(context, "index.html")
    })
}