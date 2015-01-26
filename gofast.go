// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import(
    "log"
    "fmt"
    "net/http"
    "os"
    "sort"
    "time"
)

const (
    PORT    string = ":8080"
    VERSION string = "1.0-beta"
)

type Gofast struct {
    router     *Router
    templating *Templating
}

// Bootstraps a new instance
func Bootstrap() *Gofast {
    log.Printf("gofast v%s", VERSION)

    router 	   := NewRouter()
    templating := NewTemplating()

    return &Gofast{&router, &templating}
}

// Returns Router instance
func (g *Gofast) GetRouter() *Router {
    return g.router
}

// Returns Templating instance
func (g *Gofast) GetTemplating() *Templating {
    return g.templating
}

// Handles HTTP requests
func (g *Gofast) Handle() {
    sort.Sort(RouteLen(g.GetRouter().GetRoutes()))
    http.Handle("/", g)

    assetsDirectory := g.GetTemplating().GetAssetsDirectory()
    assetsUrl       := fmt.Sprintf("/%s/", assetsDirectory)
    assetsPrefix    := fmt.Sprintf("/%s", assetsDirectory)

    http.Handle(assetsUrl, http.StripPrefix(assetsPrefix, http.FileServer(http.Dir(assetsDirectory))))

    port := PORT

    if p := os.Getenv("PORT"); p != "" {
        port = fmt.Sprintf(":%s", p)
    }

    http.ListenAndServe(port, nil)
}

// Serves HTTP request by matching the correct route
func (g *Gofast) ServeHTTP(res http.ResponseWriter, req *http.Request) {
    router  := g.GetRouter()
    fallbackRoute := router.GetFallback()

    for _, route := range router.GetRoutes() {
        if req.Method == route.method && route.pattern.MatchString(req.URL.Path) {
        	context := NewContext()
            context.SetRequest(req, route)
            context.SetResponse(res)

            startTime := time.Now()

            if (fallbackRoute.name == "fallback" && req.URL.Path != "/" && route.pattern.String() == "/") {
                route = fallbackRoute
            }

            route.handler(context)

            stopTime := time.Now()

            log.Printf("[%s] %v | route: '%s' | url: %q (time: %v)\n", req.Method, context.GetResponse().GetStatusCode(), route.name, req.URL.String(), stopTime.Sub(startTime))
            break
        }
    }
}