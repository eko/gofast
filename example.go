// An exemple application written with Gofast
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package main

import (
    "./lib"
    "net/http"
)

func main() {
    c          := gofast.Bootstrap().GetContext()
    router     := c.GetRouter()
    templating := c.GetTemplating()

    templating.SetDirectory("views")

    router.Get("index", "/", func(res http.ResponseWriter, req *http.Request) {
        templating.Render(c, "index.html")
    })

    router.Get("toto", "/toto/[0-9]+", func(res http.ResponseWriter, req *http.Request) {
        templating.Render(c, "toto.html")
    })

    router.Post("post", "/post", func(res http.ResponseWriter, req *http.Request) {
        templating.Render(c, "post.html")
    })

    c.Handle()
}
