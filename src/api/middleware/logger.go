package middleware

import (
	"net/http"
	"log"
	"time"
)

func Logger(handler http.Handler) http.Handler {
	return aliceWrapper(loggerMiddleware)(handler)
}

func loggerMiddleware(rw http.ResponseWriter, r *http.Request, next http.Handler) {
	t1 := time.Now()
	next.ServeHTTP(rw, r)
	log.Printf("\t%s | [%s] %q | %v\n", r.Host, r.Method, r.URL.String(), time.Now().Sub(t1))
}
