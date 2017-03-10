Gofast - A light, fast and simple Go micro-framework
====================================================

[![GoDoc](https://godoc.org/github.com/eko/gofast?status.png)](https://godoc.org/github.com/eko/gofast)
[![Build Status](https://secure.travis-ci.org/eko/gofast.png?branch=master)](http://travis-ci.org/eko/gofast)

This is a light and fast micro-framework I wrote in order to train to Go language.

This project uses "pongo2" library for rendering templates (compatible with Django Jinja templates)

Installation
------------

```bash
$ git clone git@github.com:eko/gofast.git
$ go get -u github.com/flosch/pongo2
$ go get -u golang.org/x/net/http2
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
    app.Get("homepage", "/$", func(context gofast.Context) {
        app.Render(context, "index.html")
    })

    // ... or add a more complex POST route with a URL parameter
    app.Post("add", "/add/([a-zA-Z]+)$", func(context gofast.Context) {
        request  := context.GetRequest()

        pattern := context.GetRoute().GetPattern()
        url     := request.GetHttpRequest().URL.Path

        request.AddParameter("name", pattern.FindStringSubmatch(url)[1])

        // ... your custom code

        app.Render(context, "add.html")
    })

    app.Listen()
}
```

HTTP/2 Support
--------------

You can use HTTP/2 support by using the following line instead of app.Listen():

```
app.ListenHttp2("./fullchain.pem", "./privkey.pem")
```

Of course, you will have to precize SSL certificate and private key.

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

Middlewares
-----------

You can add some middlewares in your application by the following way:

```go
app.Use(func(context gofast.Context, next gofast.MiddlewareFunc) gofast.Handler {
  // Some code before calling the next middleware
  handler := next(context, next)
  // Some code after calling the next middleware

  return handler
})
```

It allows you to access `context` (request, response, current route) and also
allows to define a new `handler` function to update the application behavior.

Default CORS headers
--------------------

Following CORS headers are enabled by default when the request has an "Origin" header:

```
Access-Control-Allow-Headers: Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization
Access-Control-Allow-Methods: POST, GET, OPTIONS, PUT, DELETE
Access-Control-Allow-Origin: domainname.tld
```

You can override any header (including CORS ones) by the following way into each action:

```go
app.Get("retrieve-data", "/retrieve$", func(context gofast.Context) {
    response := context.GetResponse()
    response.Header().Set("Access-Control-Allow-Methods", "GET")
    response.Header().Set("Content-Type", "application/json")

    fmt.Fprint(response, "{result: 200}")
})
```
