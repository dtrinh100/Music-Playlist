package middleware

import (
	"github.com/dtrinh100/Music-Playlist/src/api/usecases"
	"net/http"
)

type MiddlewareHandler struct {
	Logger usecases.Logger
	next   http.Handler
}
