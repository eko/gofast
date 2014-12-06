// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
    "net/http"
    "testing"
)

// Tests different add methods
func TestAllAddMethods(t *testing.T) {
    router := NewRouter()

    router.Get("get", "/get", func(res http.ResponseWriter, req *http.Request) {})
    router.Post("post", "/post", func(res http.ResponseWriter, req *http.Request) {})
    router.Patch("patch", "/patch", func(res http.ResponseWriter, req *http.Request) {})
    router.Put("put", "/put", func(res http.ResponseWriter, req *http.Request) {})
    router.Delete("delete", "/delete", func(res http.ResponseWriter, req *http.Request) {})
    router.Options("options", "/options", func(res http.ResponseWriter, req *http.Request) {})
    router.Head("head", "/head", func(res http.ResponseWriter, req *http.Request) {})

    if (7 != len(router.GetRoutes())) {
        t.Fail()
    }

    route := router.GetRoute("get")
    if (route.method != "GET" || route.pattern.String() != "/get") { t.Fail() }

    route = router.GetRoute("post")
    if (route.method != "POST" || route.pattern.String() != "/post") { t.Fail() }

    route = router.GetRoute("patch")
    if (route.method != "PATCH" || route.pattern.String() != "/patch") { t.Fail() }

    route = router.GetRoute("put")
    if (route.method != "PUT" || route.pattern.String() != "/put") { t.Fail() }

    route = router.GetRoute("delete")
    if (route.method != "DELETE" || route.pattern.String() != "/delete") { t.Fail() }

    route = router.GetRoute("options")
    if (route.method != "OPTIONS" || route.pattern.String() != "/options") { t.Fail() }

    route = router.GetRoute("head")
    if (route.method != "HEAD" || route.pattern.String() != "/head") { t.Fail() }
}
