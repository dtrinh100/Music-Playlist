package middleware

import (
	"net/http"
	"github.com/justinas/alice"
	"bytes"
	"log"
	"os"
)

type middlewareSig func(rw http.ResponseWriter, r *http.Request, next http.Handler)

/**
	This wrapper-function is used to create clean-looking middleware-functions
	for justinas/alice.
*/
func aliceWrapper(middleware middlewareSig) alice.Constructor {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			middleware(rw, r, next)
		})
	}
}

/**
	captureOutputExpectResponse captures log's output to a buffer and returns
	the output + an *http.Response. This will be used to test methods that
	output to the log.
*/
func captureOutputExpectResponse(f (func() *http.Response)) (*http.Response, string) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	res := f()
	log.SetOutput(os.Stderr)
	return res, buf.String()
}
