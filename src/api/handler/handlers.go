package handler

import (
	"github.com/dtrinh100/Music-Playlist/src/api/common"
	"log"
	"net/http"
)

/**
HandlerEnvFn is a signature for a function that includes http.Handler + env
parameters & returns an error-type
*/
type HandlerEnvFn func(rw http.ResponseWriter, req *http.Request, env *common.Env) error

/**
Handler struct takes a configured Env and a function matching our
useful signature.
*/
type Handler struct {
	*common.Env
	HEF HandlerEnvFn
}

/**
ServeHTTP allows our Handler type to satisfy http.Handler.
*/
func (h Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	handleErrorFn := func(err error) {
		switch e := err.(type) {
		case JSONError:
			common.JSONResponse(rw, e.DataMap, e.Status())
		default:
			log.Println("Custom Error-type needs to be handled in switch-statement.")
			common.JSONErrorResponse(rw)
		}
	}

	handlerErr := h.HEF(rw, req, h.Env)
	if handlerErr != nil {
		switch e := handlerErr.(type) {
		case Error:
			log.Printf("HTTP %d - %s", e.Status(), e)
			handleErrorFn(handlerErr)
		default:
			common.JSONErrorResponse(rw)
		}
	}
}
