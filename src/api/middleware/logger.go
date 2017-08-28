package middleware

import (
	"net/http"
	"log"
	"time"
)

/**
	This function is used as the public middleware-function.
*/
func Logger(handler http.Handler) http.Handler {
	return aliceWrapper(loggerMiddleware)(handler)
}

/**
	This middleware-function logs server-request information to the terminal.
*/
func loggerMiddleware(rw http.ResponseWriter, r *http.Request, next http.Handler) {
	t1 := time.Now()
	next.ServeHTTP(rw, r)
	log.Printf("\t%s | [%s] %q | %v\n", r.Host, r.Method, r.URL.String(), time.Now().Sub(t1))
}
