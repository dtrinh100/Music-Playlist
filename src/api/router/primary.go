package router

import (
	mw "github.com/dtrinh100/Music-Playlist/src/api/middleware"
	"github.com/dtrinh100/Music-Playlist/src/api/common"
	gmux "github.com/gorilla/mux"
	"github.com/dtrinh100/Music-Playlist/src/api/handler"
	"github.com/justinas/alice"

	"net/http"
)

func HandleFn(fn handler.Mp_env_fn, env *common.Env) handler.Handler {
	return handler.Handler{
		Env: env,
		H:   fn,
	}
}

/**
	This function sets-up the routes & middleware.
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

	// Global middleware(s)
	globalMiddlewares := []alice.Constructor{
		mw.Logger,
	}

	return alice.New(globalMiddlewares...).Then(router)
}
