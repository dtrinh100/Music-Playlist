package middleware

import (
	"github.com/dtrinh100/Music-Playlist/src/api/usecases"
	"github.com/dtrinh100/Music-Playlist/src/api/interfaces"
	"net/http"
)

type MiddlewareHandler struct {
	Responder interfaces.WebResponder
	Logger    usecases.Logger
	next      http.Handler
}
