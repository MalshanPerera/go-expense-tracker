package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := range xs {
			x := xs[i]
			next = x(next)
		}

		return next
	}
}
