// An exemple application written with Gofast
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package main

import (
    "./lib"
    "net/http"
    "fmt"
    "time"
    "log"
)

func main() {
    g      := gofast.Bootstrap()
    router := g.GetRouter()

    router.Add("index", "/", func(w http.ResponseWriter, r *http.Request) {
        t1 := time.Now()
        fmt.Fprintf(w, "<h1>Welcome!</h1>")
        t2 := time.Now()
        log.Printf("[%s] %q (time: %v)\n", r.Method, r.URL.String(), t2.Sub(t1))
    })

    router.Add("toto", "/toto", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h1>Toto webpage!</h1>")
    })

    g.Handle()
}
