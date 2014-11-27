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

    router.Get("index", "/", func(w http.ResponseWriter, r *http.Request) {
        templating.Render(w, "index.html")
    })

    router.Get("toto", "/toto", func(w http.ResponseWriter, r *http.Request) {
        templating.Render(w, "toto.html")
    })

    router.Post("post", "/post", func(w http.ResponseWriter, r *http.Request) {
        templating.Render(w, "post.html")
    })

    g.Handle()
}
