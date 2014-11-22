// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import(
    "net/http"
    "os"
    "fmt"
    "log"
    "time"
)

const (
    VERSION string = "1.0-beta"
)

type gofast struct {
    logger *log.Logger
    router router
}

// Bootstraps a new instance
func Bootstrap() gofast {
    log.Printf("gofast v%s", VERSION)

    logger := log.New(os.Stdout, "[gofast]", 0)
    router := NewRouter()

    return gofast{logger: logger, router: router}
}

// Handles HTTP requests
func (g *gofast) Handle() {
    http.HandleFunc("/", indexHandler)
    http.ListenAndServe(":8080", nil)
}

// Returns router
func (g *gofast) Router() router {
    return g.router
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    t1 := time.Now()
    fmt.Fprintf(w, "<h1>Welcome!</h1>")
    t2 := time.Now()
    log.Printf("[%s] %q (time: %v)\n", r.Method, r.URL.String(), t2.Sub(t1))
}