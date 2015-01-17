Gofast - A simple Go micro-framework
====================================

[![Build Status](https://secure.travis-ci.org/eko/gofast.png?branch=master)](http://travis-ci.org/eko/gofast)

This is a micro-framework I wrote in order to train to Golang language.

This project uses "pongo2" library for rendering templates (compatible with Django Jinja templates)

Installation
------------

```bash
$ git clone git@github.com:eko/gofast.git
$ go get -u github.com/flosch/pongo2
```

Running an application
----------------------

```bash
$ go run app.go
2014/12/06 15:40:28 gofast v1.0-beta
2014/12/06 15:40:32 [GET] 200 "/" (time: 143.238us)
```

A simple application example
----------------------------

Because an example will explain it better, here is an application example with things you can do with Gofast:

```go
package main

import (
    "github.com/eko/gofast"
)

func main() {
    c          := gofast.Bootstrap().GetContext()
    router     := c.GetRouter()
    templating := c.GetTemplating()

    templating.SetAssetsDirectory("assets")
    templating.SetViewsDirectory("views")

    // This add a fallback route for 404 (not found) resources
    router.SetFallback(func() {
        templating.Render(c, "404.html")
    })

    // You can add a simple GET route
    router.Get("homepage", "/", func() {
        templating.Render(c, "index.html")
    })

    // ... or add a more complex POST route with a URL parameter
    router.Post("add", "/add/([0-9]+)", func() {
        request  := c.GetRequest()

        pattern := request.GetRoute().GetPattern()
        url     := request.GetHttpRequest().URL.Path

        request.AddParameter("name", pattern.FindStringSubmatch(url)[1])

        // ... your custom code

        templating.Render(c, "add.html")
    })

    // If you call the /test418 URL, a 418 (I'm a teapot) status code will be rendered
    router.Get("test418", "/test418", func() {
        c.GetResponse().SetStatusCode(418)
    })

    c.Handle()
}
```

Requesting this example
-----------------------

Using the example given below, here is the request results:

```bash
> $ curl -X GET http://127.0.0.1:8080/
<h1>Welcome to the index template!</h1>

> $ curl -X POST http://127.0.0.1:8080/add/toto
<h1>Added: toto</h1>
```
