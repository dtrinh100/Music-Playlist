package router

import (
	"github.com/dtrinh100/Music-Playlist/src/api/common"
	"github.com/dtrinh100/Music-Playlist/src/api/handler"
	mw "github.com/dtrinh100/Music-Playlist/src/api/middleware"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

/**
SetAuthRoutes sets-up the '/api/auth' routes.
*/
func SetAuthRoutes(router *mux.Router, env *common.Env) *mux.Router {
	router.Handle("/register", HandleFn(handler.Register, env)).Methods("POST")
	router.Handle("/login", HandleFn(handler.Login, env)).Methods("POST")

	return router
}
