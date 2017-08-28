package middleware

import (
	"net/http"
	"github.com/justinas/alice"
)

type middlewareSig func(rw http.ResponseWriter, r *http.Request, next http.Handler)

func aliceWrapper(middleware middlewareSig) alice.Constructor {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			middleware(rw, r, next)
		})
	}
}
