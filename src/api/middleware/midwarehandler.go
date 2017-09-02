package middleware

import (
	"github.com/dtrinh100/Music-Playlist/src/api/common"
	"net/http"
)

/**
	Mp_aliceenv_fn is a signature-wrapper that will allow support for this type of signature
	to work as justinas/alice middleware while providing the '*common.Env' struct.
**/
type Mp_aliceenv_fn func(rw http.ResponseWriter, req *http.Request, next http.Handler, env *common.Env)

/**
	AliceMiddlewareHandler helps functions with the signature Mp_aliceenv_fn to be passed
	into justinas/alice's functions.
 */
type AliceMiddlewareEnvHandler struct {
	*common.Env
	AliceFn Mp_aliceenv_fn
}

// Handle mimics the same signature that justina/alice's Constructor requires.
func (ameh AliceMiddlewareEnvHandler) Handle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ameh.AliceFn(rw, req, h, ameh.Env)
	})
}

/**
	mp_alice_fn is a signature-wrapper that will allow support for this type of signature
	to function as justinas/alice middleware.
**/
type Mp_alice_fn func(rw http.ResponseWriter, req *http.Request, next http.Handler)

/**
	AliceMiddlewareHandler helps functions with the signature Mp_alice_fn to be passed into
	justinas/alice's functions.
 */
type AliceMiddlewareHandler struct {
	AliceFn Mp_alice_fn
}

// Handle mimics the same signature that justina/alice's Constructor requires.
func (amh AliceMiddlewareHandler) Handle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		amh.AliceFn(rw, req, h)
	})
}
