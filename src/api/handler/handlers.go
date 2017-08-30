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

// The Handler struct that takes a configured Env and a function matching
// our useful signature.
type Handler struct {
	*Env
	H func(e *Env, w http.ResponseWriter, r *http.Request) error
}

// ServeHTTP allows our Handler type to satisfy http.Handler.
func (h Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	err := h.H(h.Env, rw, req)
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
