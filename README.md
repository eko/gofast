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

Example
-------

You can find an application example in `example.go` file and run it simply by typing:

```bash
$ go run example.go
2014/12/06 15:40:28 gofast v1.0-beta
2014/12/06 15:40:32 [GET] 200 "/" (time: 143.238us)
```

Running example
---------------

Using the `example.go` file given in this repository, here is the result:

```bash
> $ curl -X GET http://127.0.0.1:8080/
<h1>Welcome to the index template!</h1>

<strong>Route name</strong>: index

> $ curl -X GET http://127.0.0.1:8080/toto/1
<h1>Toto #1</h1>

<strong>Route name</strong>: toto

> $ curl -X POST http://127.0.0.1:8080/post
POST method handled!

<strong>Route name</strong>: post
```