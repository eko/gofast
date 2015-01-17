// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
    "testing"
)

// Tests different add methods
func TestAllAddMethods(t *testing.T) {
    router := NewRouter()

    router.Get("get", "/get", func() {})
    router.Post("post", "/post", func() {})
    router.Patch("patch", "/patch", func() {})
    router.Put("put", "/put", func() {})
    router.Delete("delete", "/delete", func() {})
    router.Options("options", "/options", func() {})
    router.Head("head", "/head", func() {})

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

// Tests adding a fallback route
func TestFallbackRoute(t *testing.T) {
    router := NewRouter()

    router.SetFallback(func() {})

    fallback := router.GetFallback()

    if "fallback" != fallback.name {
        t.Fail()
    }
}