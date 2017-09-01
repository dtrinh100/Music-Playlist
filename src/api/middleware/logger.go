package middleware

import (
	"net/http"
	"log"
	"time"
)

/**
	LoggerMiddleware is a function that logs server-request information to the terminal.
*/
func LoggerMiddleware(rw http.ResponseWriter, req *http.Request, next http.Handler) {
	t1 := time.Now()
	next.ServeHTTP(rw, req)
	log.Printf("\t%s | [%s] %q | %v\n", req.Host, req.Method, req.URL.String(), time.Now().Sub(t1))
}
