package router

import (
	gmux "github.com/gorilla/mux"
	"github.com/dtrinh100/Music-Playlist/src/api/handler"
)

func InitializeRoutes(env *handler.Env) *gmux.Router {
	// Using Gorilla mux router instead of default one because it offers more flexibility
	router := gmux.NewRouter().StrictSlash(false)

	// API sub-router
	// NOTE: every route has to go through '/api' as of now.
	//		This is due to the way things are configured w/ Docker
	apiRouter := router.PathPrefix("/api").Subrouter()

	// User sub-router & routes
	userRouter := apiRouter.PathPrefix("/users").Subrouter()
	userRouter = SetUserRoutes(env, userRouter)

	// Auth sub-router & routes
	authRouter := apiRouter.PathPrefix("/auth").Subrouter()
	authRouter = SetAuthRoutes(authRouter)

	return router
}
