package router

import (
	"github.com/gorilla/mux"

	"github.com/dtrinh100/Music-Playlist/src/api/handler"
)

/**
	This function sets up the '/users' routes
*/
func SetUserRoutes(router *mux.Router, eh *handler.EnvHandler) *mux.Router {
	// route handlers
	// ignore vet errors for unkeyed fields
	router.Handle("", eh.Handle(handler.PostUser)).Methods("POST")

	return router
}

// TODO: fill out code-block
func SetAuthenticatedUserRoutes(router *mux.Router) *mux.Router {
	return nil
}
