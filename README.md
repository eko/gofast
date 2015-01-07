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

    router.Get("homepage", "/", func() {
        templating.Render(c, "index.html")
    })

    router.Post("add", "/add/([0-9]+)", func() {
        request  := c.GetRequest()

        pattern := request.GetRoute().GetPattern()
        url     := request.GetHttpRequest().URL.Path

        request.AddParameter("name", pattern.FindStringSubmatch(url)[1])

        // ... your custom code

        templating.Render(c, "add.html")
    })

    router.Get("test404", "/test404", func() {
        c.GetResponse().SetStatusCode(404)
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
