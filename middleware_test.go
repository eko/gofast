// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

import (
	"testing"
)

// Tests initializing a new middleware component
func TestMiddleware(t *testing.T) {
	middleware := NewMiddleware()

	if len(middleware.middlewares) > 0 {
		t.Fail()
	}
}

// Tests adding a new middlewares
func TestUseNewMiddlewares(t *testing.T) {
	middleware := NewMiddleware()
	middleware.Use(func(context Context, next MiddlewareFunc) Handler {
		return next(context, next)
	})

	middleware.Use(func(context Context, next MiddlewareFunc) Handler {
		return next(context, next)
	})

	if len(middleware.middlewares) < 2 {
		t.Fail()
	}
}
