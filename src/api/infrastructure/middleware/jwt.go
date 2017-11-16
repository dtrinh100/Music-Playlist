package middleware

import (
	"net/http"
)

type JWTHandler struct {
	MiddlewareHandler
}

// TODO: implement
func (handler *JWTHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	handler.Logger.Log("Called JWT middleware")
	handler.next.ServeHTTP(rw, req)
}

func (handler *JWTHandler) Handle(next http.Handler) http.Handler {
	handler.next = next
	return handler
}
