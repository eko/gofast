// An exemple application written with Gofast
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package main

import (
    "./lib"
    "net/http"
)

func main() {
    g          := gofast.Bootstrap()
    router     := g.GetRouter()
    templating := g.GetTemplating()

    templating.SetDirectory("views")

    router.Get("index", "/", func(res http.ResponseWriter, req *http.Request) {
        templating.Render(res, "index.html")
    })

    router.Get("toto", "/toto", func(res http.ResponseWriter, req *http.Request) {
        templating.Render(res, "toto.html")
    })

    router.Post("post", "/post", func(res http.ResponseWriter, req *http.Request) {
        templating.Render(res, "post.html")
    })

    g.Handle()
}
