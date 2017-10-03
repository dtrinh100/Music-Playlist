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

/**
SetAuthenticatedAuthRoutes sets-up the '/api/auth' authenticated-routes.
*/
func SetAuthenticatedAuthRoutes(router *mux.Router, env *common.Env) *mux.Router {
	authMiddleware := alice.New(
		AliceEnv(mw.JWTMiddleware, env).Handle,
	)

	router.Handle("/verify", authMiddleware.Then(
		HandleFn(handler.Verify, env))).Methods("GET")
	router.Handle("/logout", authMiddleware.Then(
		HandleFn(handler.Logout, env))).Methods("GET")

	return router
}
