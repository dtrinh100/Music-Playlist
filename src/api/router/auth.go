package router

import (
	"github.com/gorilla/mux"
	"github.com/dtrinh100/Music-Playlist/src/api/handler"
	"github.com/dtrinh100/Music-Playlist/src/api/common"
)

/**
	SetAuthRoutes sets-up the '/api/auth' routes.
*/
func SetAuthRoutes(router *mux.Router, env *common.Env) *mux.Router {
	router.Handle("", HandleFn(handler.Login, env)).Methods("POST")
	return router
}
