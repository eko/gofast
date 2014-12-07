// An exemple application written with Gofast
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package main

import (
    "./lib"
)

func main() {
    c          := gofast.Bootstrap().GetContext()
    router     := c.GetRouter()
    templating := c.GetTemplating()

    templating.SetAssetsDirectory("assets")
    templating.SetViewsDirectory("views")

    router.Get("index", "/", func() {
        templating.Render(c, "index.html")
    })

    router.Post("post", "/post", func() {
        templating.Render(c, "post.html")
    })

    router.Get("toto", "/toto/([0-9]+)", func() {
        request  := c.GetRequest()

        pattern := request.GetRoute().GetPattern()
        url     := request.GetHttpRequest().URL.Path

        request.AddParameter("identifier", pattern.FindStringSubmatch(url)[1])

        templating.Render(c, "toto.html")
    })

    router.Get("test404", "/test404", func() {
        c.GetResponse().SetStatusCode(404)
    })

    c.Handle()
}
