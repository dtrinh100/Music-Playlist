package router

import (
	"github.com/gorilla/mux"

	"github.com/dtrinh100/Music-Playlist/src/api/common"
	"github.com/dtrinh100/Music-Playlist/src/api/handler"
)

/**
	This function sets up the '/users' routes
*/
func SetUserRoutes(router *mux.Router, env *common.Env) *mux.Router {
	// route handlers
	// ignore vet errors for unkeyed fields
	router.Handle("", HandleFn(handler.PostUser, env)).Methods("POST")

	return router
}

// TODO: fill out code-block
func SetAuthenticatedUserRoutes(router *mux.Router) *mux.Router {
	return nil
}
