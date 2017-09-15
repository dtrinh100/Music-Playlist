package router

import (
	"github.com/dtrinh100/Music-Playlist/src/api/common"
	"github.com/dtrinh100/Music-Playlist/src/api/handler"
	"github.com/gorilla/mux"
)

/**
SetAuthRoutes sets-up the '/api/auth' routes.
*/
func SetAuthRoutes(router *mux.Router, env *common.Env) *mux.Router {
	router.Handle("", HandleFn(handler.Login, env)).Methods("POST")
	return router
}
