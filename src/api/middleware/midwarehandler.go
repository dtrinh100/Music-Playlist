package middleware

import (
	"github.com/dtrinh100/Music-Playlist/src/api/common"
	"net/http"
)

/**
	AliceEnvFn is a signature-wrapper that will allow support for this type of signature
	to work as justinas/alice middleware while providing the '*common.Env' struct.
**/
type AliceEnvFn func(rw http.ResponseWriter, req *http.Request, next http.Handler, env *common.Env)

/**
AliceMiddlewareEnvHandler helps functions with the signature AliceEnvFn to be passed
into justinas/alice's functions.
*/
type AliceMiddlewareEnvHandler struct {
	*common.Env
	AliceEnvFn AliceEnvFn
}

// Handle mimics the same signature that justina/alice's Constructor requires.
func (ameh AliceMiddlewareEnvHandler) Handle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ameh.AliceEnvFn(rw, req, h, ameh.Env)
	})
}

/**
	AliceFn is a signature-wrapper that will allow support for this type of signature
	to function as justinas/alice middleware.
**/
type AliceFn func(rw http.ResponseWriter, req *http.Request, next http.Handler)

/**
AliceMiddlewareHandler helps functions with the signature AliceFn to be passed into
justinas/alice's functions.
*/
type AliceMiddlewareHandler struct {
	AliceFn AliceFn
}

// Handle mimics the same signature that justina/alice's Constructor requires.
func (amh AliceMiddlewareHandler) Handle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		amh.AliceFn(rw, req, h)
	})
}
