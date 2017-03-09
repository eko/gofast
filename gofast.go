// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"time"

	"golang.org/x/net/http2"
)

const (
	PORT    string = ":8080"
	VERSION string = "1.0-beta"
)

type Gofast struct {
	*Router
	*Templating
	*Middleware
}

// Bootstraps a new instance
func Bootstrap() *Gofast {
	log.Printf("gofast v%s", VERSION)

	router := NewRouter()
	templating := NewTemplating()
	middleware := NewMiddleware()

	return &Gofast{&router, &templating, &middleware}
}

// Prepares a HTTP server
func (g *Gofast) PrepareHttpServer() string {
	sort.Sort(RouteLen(g.GetRoutes()))
	http.Handle("/", g)

	assetsDirectory := g.GetAssetsDirectory()
	assetsUrl := fmt.Sprintf("/%s/", assetsDirectory)
	assetsPrefix := fmt.Sprintf("/%s", assetsDirectory)

	http.Handle(assetsUrl, http.StripPrefix(assetsPrefix, http.FileServer(http.Dir(assetsDirectory))))

	port := PORT

	if p := os.Getenv("PORT"); p != "" {
		port = fmt.Sprintf(":%s", p)
	}

	return port
}

// Listens and handles HTTP requests
func (g *Gofast) Listen() {
	port := g.PrepareHttpServer()

	http.ListenAndServe(port, nil)
}

// Listens and handles HTTP/2 requests
func (g *Gofast) ListenHttp2(certificate string, key string) {
	port := g.PrepareHttpServer()
	server := &http.Server{Addr: port, Handler: nil}

	http2.ConfigureServer(server, nil)
	log.Fatal(server.ListenAndServeTLS(certificate, key))
}

// Serves HTTP request by matching the correct route
func (g *Gofast) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	matchedRoute := g.GetFallback()

	for _, route := range g.GetRoutes() {
		if "fallback" == route.name {
			continue
		}

		if (route.method == req.Method || "*" == route.method) && route.pattern.MatchString(req.URL.Path) {
			matchedRoute = route
			break
		}
	}

	g.HandleRoute(res, req, matchedRoute)
}

// Handles a route with the initialized context
func (g *Gofast) HandleRoute(res http.ResponseWriter, req *http.Request, route Route) {
	startTime := time.Now()

	context := NewContext()
	context.SetRoute(&route)
	context.SetRequest(req)
	context.SetResponse(res)

	handler := g.HandleMiddlewares(context)
	handler(context)

	stopTime := time.Now()

	log.Printf("[%s] %v | route: '%s' | url: %q (time: %v)\n", req.Method, context.GetResponse().GetStatusCode(), route.name, req.URL.String(), stopTime.Sub(startTime))
}
