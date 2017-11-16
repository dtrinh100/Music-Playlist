package middleware

import (
	"net/http"
	"time"
	"fmt"
)

type WebLoggerHandler struct {
	MiddlewareHandler
}

func (handler *WebLoggerHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	t1 := time.Now()
	handler.next.ServeHTTP(rw, req)
	msg := fmt.Sprintf("\t%s | [%s] %q | %v\n", req.Host, req.Method, req.URL.String(), time.Now().Sub(t1))
	handler.Logger.Log(msg)
}

func (handler *WebLoggerHandler) Handle(next http.Handler) http.Handler {
	handler.next = next
	return handler
}
