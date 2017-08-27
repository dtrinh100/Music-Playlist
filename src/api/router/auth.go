package router

import (
	"github.com/gorilla/mux"
	"github.com/dtrinh100/Music-Playlist/src/api/handler"
)

func SetAuthRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("", handler.Login).Methods("POST")
	return router
}
