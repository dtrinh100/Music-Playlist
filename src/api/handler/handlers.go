package handler

import (
	"log"
	"net/http"
	"github.com/dtrinh100/Music-Playlist/src/api/common"
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
			common.JSONErrorResponse(rw, e.ErrMap, e.Status())
			return
		case StatusError:
			http.Error(rw, e.Error(), e.Status())
		default:
			log.Println("Custom Error-type needs to be handled in switch-statement.")
			status := http.StatusInternalServerError
			http.Error(rw, http.StatusText(status), status)
		}
	}

	err := h.HEF(rw, req, h.Env)
	if err != nil {
		switch e := err.(type) {
		case Error:
			log.Printf("HTTP %d - %s", e.Status(), e)
			handleErrorFn(err)
		default:
			http.Error(rw, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}
