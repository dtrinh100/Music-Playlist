package router

import (
	"github.com/dtrinh100/Music-Playlist/src/api/common"
	"github.com/dtrinh100/Music-Playlist/src/api/handler"
	mw "github.com/dtrinh100/Music-Playlist/src/api/middleware"
	gmux "github.com/gorilla/mux"
	"github.com/justinas/alice"

	"net/http"
)

// Alice is a helper-function that makes middleware-functions satisfy alice.Constructor.
func Alice(fn mw.AliceFn) mw.AliceMiddlewareHandler {
	return mw.AliceMiddlewareHandler{AliceFn: fn}
}

/**
AliceEnv is a helper-function that makes middleware-functions satisfy alice.Constructor
and pass an Env struct.
*/
func AliceEnv(fn mw.AliceEnvFn, env *common.Env) mw.AliceMiddlewareEnvHandler {
	return mw.AliceMiddlewareEnvHandler{Env: env, AliceEnvFn: fn}
}

// HandleFn is a helper-function that makes route-functions satisfy http.Handler.
func HandleFn(fn handler.HandlerEnvFn, env *common.Env) handler.Handler {
	return handler.Handler{
		Env: env,
		HEF: fn,
	}
}

/**
InitializeRoutes sets-up all of the routes & middleware.
*/
func InitializeRoutes(env *common.Env) http.Handler {
	// Using Gorilla mux router instead of default one because it offers more flexibility
	router := gmux.NewRouter().StrictSlash(false)

	// API sub-router
	// NOTE: every route has to go through '/api' as of now.
	//		This is due to the way things are configured w/ Docker
	apiRouter := router.PathPrefix("/api").Subrouter()

	// User sub-router & routes
	userRouter := apiRouter.PathPrefix("/users").Subrouter()
	userRouter = SetUserRoutes(userRouter, env)

	// Auth sub-router & routes
	authRouter := apiRouter.PathPrefix("/auth").Subrouter()
	authRouter = SetAuthRoutes(authRouter, env)
	authRouter = SetAuthenticatedAuthRoutes(authRouter, env)

	// Global middleware(s)
	globalMiddlewares := []alice.Constructor{
		Alice(mw.LoggerMiddleware).Handle,
	}

	return alice.New(globalMiddlewares...).Then(router)
}
