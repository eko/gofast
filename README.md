Gofast - A simple Go micro-framework
====================================

[![GoDoc](https://godoc.org/github.com/eko/gofast?status.png)](https://godoc.org/github.com/eko/gofast)
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
2015/01/26 21:57:35 gofast v1.0-beta
2015/01/26 21:57:48 [POST] 200 | route: 'add' | url: "/add/toto" (time: 143.238us)
```

This will run the application on port 8080. Optionnaly, you can provide a port number this way:

```bash
$ PORT=8005 go run app.go
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
    g          := gofast.Bootstrap()
    router     := g.GetRouter()
    templating := g.GetTemplating()

    templating.SetAssetsDirectory("assets")
    templating.SetViewsDirectory("views")

    // This add a fallback route for 404 (not found) resources
    router.SetFallback(func(c gofast.Context) {
        c.GetResponse().SetStatusCode(404)
        templating.Render(c, "404.html")
    })

    // You can add a simple GET route
    router.Get("homepage", "/", func(c gofast.Context) {
        templating.Render(c, "index.html")
    })

    // ... or add a more complex POST route with a URL parameter
    router.Post("add", "/add/([a-zA-Z]+)", func(c gofast.Context) {
        request  := c.GetRequest()

        pattern := request.GetRoute().GetPattern()
        url     := request.GetHttpRequest().URL.Path

        request.AddParameter("name", pattern.FindStringSubmatch(url)[1])

        // ... your custom code

        templating.Render(c, "add.html")
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
