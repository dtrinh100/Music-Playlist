package middleware

import (
	"net/http"
	"bytes"
	"log"
	"os"
)

/**
	AliceFunc is a signature-wrapper that will allow support for this type of signature
	to function as justinas/alice middleware.
**/
type AliceFunc func(rw http.ResponseWriter, r *http.Request, next http.Handler)

// Handle mimics the same signature that justina/alice's Constructor requires.
func (afn AliceFunc) Handle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		afn(rw, r, h)
	})
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
