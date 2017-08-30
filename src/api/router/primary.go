package router

import (
	mw "github.com/dtrinh100/Music-Playlist/src/api/middleware"
	gmux "github.com/gorilla/mux"
	"github.com/dtrinh100/Music-Playlist/src/api/handler"
	"github.com/justinas/alice"

	"net/http"
)

/**
	This function sets-up the routes & middleware.
*/
func InitializeRoutes(eh *handler.EnvHandler) http.Handler {
	// Using Gorilla mux router instead of default one because it offers more flexibility
	router := gmux.NewRouter().StrictSlash(false)

	// API sub-router
	// NOTE: every route has to go through '/api' as of now.
	//		This is due to the way things are configured w/ Docker
	apiRouter := router.PathPrefix("/api").Subrouter()

	// User sub-router & routes
	userRouter := apiRouter.PathPrefix("/users").Subrouter()
	userRouter = SetUserRoutes(userRouter, eh)

	// Auth sub-router & routes
	authRouter := apiRouter.PathPrefix("/auth").Subrouter()
	authRouter = SetAuthRoutes(authRouter)

	// Global middleware(s)
	globalMiddlewares := []alice.Constructor{
		mw.Logger,
	}

	return alice.New(globalMiddlewares...).Then(router)
}
