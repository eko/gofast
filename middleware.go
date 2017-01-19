// A fast web framework written in Go.
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package gofast

type Middleware struct {
	middlewares []MiddlewareFunc
}

type MiddlewareFunc func(context Context, middleware MiddlewareFunc) Handler

// Creates a new middleware component instance
func NewMiddleware() Middleware {
	return Middleware{middlewares: make([]MiddlewareFunc, 0)}
}

// Adds a new middleware
func (m *Middleware) Use(middleware MiddlewareFunc) {
	m.middlewares = append(m.middlewares, middleware)
}

// Handle middlewares and returns handler
func (m *Middleware) HandleMiddlewares(context Context) Handler {
	m.Use(func(context Context, next MiddlewareFunc) Handler {
		return context.GetRoute().GetHandler()
	})

	handler := context.GetRoute().GetHandler()

	for i := 0; i < len(m.middlewares)-1; i++ {
		middleware := m.middlewares[i]
		next := m.middlewares[i+1]

		handler = middleware(context, next)
		context.GetRoute().SetHandler(handler)
	}

	return handler
}
