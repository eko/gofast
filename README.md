Gofast - A simple Go micro-framework
====================================

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
2014/11/22 12:24:54 gofast v1.0-beta
2014/11/22 12:24:55 [GET] "/" (time: 11.188us)
2014/11/22 12:24:55 [GET] "/favicon.ico" (time: 4.636us)
```

Running example
---------------

Using the `example.go` file given in this repository, here is the result:

```bash
> $ curl -X GET http://127.0.0.1:8080/
Welcome to the index template!

> $ curl -X GET http://127.0.0.1:8080/toto
0 + 0 = Toto

> $ curl -X POST http://127.0.0.1:8080/post
POST method handled!
```