package handler

import (
	"log"
	"net/http"
)

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

// Error gets the error string of the error
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Status gets the status code of the error
func (se StatusError) Status() int {
	return se.Code
}

// Signature for a function that includes http.Handler + env parameters & returns an error-type
type mp_envhandler_fn func(rw http.ResponseWriter, req *http.Request, env *Env) error

// EnvHandler represents a handler of Handler-handler
type EnvHandler struct {
	*Env
}

// Handle allows EnvHandler to create Handler handlers while reusing the *Env variable
func (eh EnvHandler) Handle(fn mp_envhandler_fn) Handler {
	return Handler{
		eh.Env,
		fn,
	}
}

// The Handler struct that takes a configured Env and a function matching
// our useful signature.
type Handler struct {
	*Env
	H mp_envhandler_fn
}

// ServeHTTP allows our Handler type to satisfy http.Handler.
func (h Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	err := h.H(rw, req, h.Env)
	if err != nil {
		switch e := err.(type) {
		case Error:
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(rw, e.Error(), e.Status())
		default:
			http.Error(rw, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}
