package router

import (
	"github.com/gorilla/mux"
	"github.com/dtrinh100/Music-Playlist/src/api/handler"
)

/**
	This function sets-up the '/api/auth' routes.
*/
func SetAuthRoutes(router *mux.Router, eh *handler.EnvHandler) *mux.Router {
	router.Handle("", eh.Handle(handler.Login)).Methods("POST")
	return router
}
