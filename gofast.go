// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
	"fmt"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
)

const (
	PORT    string = ":8080"
	VERSION string = "1.0-beta"
)

type Gofast struct {
	*logrus.Logger
	*Router
	*Templating
	*Middleware
}

// Bootstrap bootstraps a new instance
func Bootstrap() *Gofast {
	logrus.WithFields(logrus.Fields{"version": VERSION}).Info("gofast is running")

	logger := logrus.New()
	router := NewRouter()
	templating := NewTemplating()
	middleware := NewMiddleware()

	return &Gofast{logger, &router, &templating, &middleware}
}

// PrepareHttpServer prepares a HTTP server
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

// Listen listens and handles HTTP requests
func (g *Gofast) Listen() {
	port := g.PrepareHttpServer()

	http.ListenAndServe(port, nil)
}

// ListenHttp2 listens and handles HTTP/2 requests
func (g *Gofast) ListenHttp2(certificate string, key string) {
	port := g.PrepareHttpServer()
	server := &http.Server{Addr: port, Handler: nil}

	http2.ConfigureServer(server, nil)
	logrus.Fatal(server.ListenAndServeTLS(certificate, key))
}

// ServeHTTP serves HTTP request by matching the correct route
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

// HandleRoute handles a route with the initialized context
func (g *Gofast) HandleRoute(res http.ResponseWriter, req *http.Request, route Route) {
	startTime := time.Now()

	context := NewContext()
	context.SetLogger(g.Logger)
	context.SetRoute(&route)
	context.SetRequest(req)
	context.SetResponse(res)

	handler := g.HandleMiddlewares(context)
	handler(context)

	stopTime := time.Now()

	logrus.WithFields(logrus.Fields{
		"name":     route.name,
		"method":   req.Method,
		"code":     context.GetResponse().GetStatusCode(),
		"url":      req.URL.String(),
		"duration": stopTime.Sub(startTime),
	}).Info("Route matched")
}
