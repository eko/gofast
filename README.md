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
    app := gofast.Bootstrap()
    app.SetAssetsDirectory("assets")
    app.SetViewsDirectory("views")

    // This adds a fallback route for 404 (not found) resources
    app.SetFallback(func(context gofast.Context) {
        context.GetResponse().SetStatusCode(404)
        app.Render(context, "404.html")
    })

    // You can add a simple GET route
    app.Get("homepage", "/", func(context gofast.Context) {
        app.Render(context, "index.html")
    })

    // ... or add a more complex POST route with a URL parameter
    app.Post("add", "/add/([a-zA-Z]+)", func(context gofast.Context) {
        request  := context.GetRequest()

        pattern := request.GetRoute().GetPattern()
        url     := request.GetHttpRequest().URL.Path

        request.AddParameter("name", pattern.FindStringSubmatch(url)[1])

        data := request.getFormValue()

        // ... your custom code

        app.Render(context, "add.html")
    })

    app.Listen()
}
```

Templating
----------

You can use all Pongo2 template features and retrieve data like this:

```twig
{% extends "../layout.html" %}

{% block navigation %}
    {% include "../include/navigation.html" with current="blog" %}
{% endblock %}

<h2>Retrieve a "name" parameter</h2>
<p>{{ request.GetParameter("name") }}</p>

<h2>Retrieve a "data" POST form value</h2>
<p>{{ request.GetFormValue("data") }}</p>
```

You have access to both `request` and `response` objects from context.

Requesting this example
-----------------------

Using the example given below, here is the request results:

```bash
> $ curl -X GET http://127.0.0.1:8080/
<h1>Welcome to the index template!</h1>

> $ curl -X POST -d'data=my post data' http://127.0.0.1:8080/add/toto
<h2>Retrieve a "name" parameter</h2>
<p>toto</p>

<h2>Retrieve a "data" POST form value</h2>
<p>my post data</p>
```
