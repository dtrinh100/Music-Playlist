package infrastructure

import (
	"github.com/dtrinh100/Music-Playlist/src/api/infrastructure/middleware"
	"github.com/dtrinh100/Music-Playlist/src/api/interfaces"
	"github.com/justinas/alice"
	gmux "github.com/gorilla/mux"

	"net/http"
)

func initUserRoutes(authRouter *gmux.Router, webservice *interfaces.WebserviceHandler) {
	userRouter := authRouter.PathPrefix("/api/auth/users").Subrouter()
	usernameRegex := "{username:[a-zA-Z0-9]+}"

	userRouter.HandleFunc("", func(rw http.ResponseWriter, req *http.Request) {
		webservice.Users(rw, req)
	}).Methods("GET")

	userRouter.HandleFunc("/"+usernameRegex, func(rw http.ResponseWriter, req *http.Request) {
		webservice.User(rw, req)
	}).Methods("GET")

	userRouter.HandleFunc("/"+usernameRegex, func(rw http.ResponseWriter, req *http.Request) {
		webservice.UpdateUser(rw, req)
	}).Methods("PATCH")

	userRouter.HandleFunc("/"+usernameRegex, func(rw http.ResponseWriter, req *http.Request) {
		webservice.DeleteUser(rw, req)
	}).Methods("DELETE")
}

func initSongRoutes(authRouter *gmux.Router, webservice *interfaces.WebserviceHandler) {
	songRouter := authRouter.PathPrefix("/api/auth/songs").Subrouter()
	songRegex := "{id:[0-9]+}"

	songRouter.HandleFunc("", func(rw http.ResponseWriter, req *http.Request) {
		webservice.Songs(rw, req)
	}).Methods("GET")

	songRouter.HandleFunc("/"+songRegex, func(rw http.ResponseWriter, req *http.Request) {
		webservice.Song(rw, req)
	}).Methods("GET")

	songRouter.HandleFunc("", func(rw http.ResponseWriter, req *http.Request) {
		webservice.CreateSong(rw, req)
	}).Methods("POST")

	songRouter.HandleFunc("/"+songRegex, func(rw http.ResponseWriter, req *http.Request) {
		webservice.UpdateSong(rw, req)
	}).Methods("PATCH")

	songRouter.HandleFunc("/"+songRegex, func(rw http.ResponseWriter, req *http.Request) {
		webservice.DeleteSong(rw, req)
	}).Methods("DELETE")
}

func initAuthRoutes(apiRouter *gmux.Router, webservice *interfaces.WebserviceHandler, jwt *middleware.JWTMiddleware) {
	authRouter := gmux.NewRouter()
	authHandle := alice.New(jwt.Handle).Then(authRouter)
	// Note: the following 2 lines are needed to allow sub-routers to be used as a handler
	apiRouter.Handle("/auth", authHandle)
	apiRouter.Handle("/auth/{path:.*}", authHandle)

	authRouter.HandleFunc("/api/auth/verify", func(rw http.ResponseWriter, req *http.Request) {
		webservice.VerifyUser(rw, req)
	})

	authRouter.HandleFunc("/api/auth/logout", func(rw http.ResponseWriter, req *http.Request) {
		webservice.LogoutUser(rw, req)
	})

	initUserRoutes(authRouter, webservice)
	initSongRoutes(authRouter, webservice)
}

func initUnauthRoutes(apiRouter *gmux.Router, webservice *interfaces.WebserviceHandler) {
	// TODO: figure out if there's a RESTful way of doing authentication
	apiRouter.HandleFunc("/register", func(rw http.ResponseWriter, req *http.Request) {
		webservice.RegisterUser(rw, req)
	}).Methods("POST")

	apiRouter.HandleFunc("/login", func(rw http.ResponseWriter, req *http.Request) {
		webservice.LoginUser(rw, req)
	}).Methods("POST")

}

func GetRouterWithRoutes(webservice *interfaces.WebserviceHandler, jwt *middleware.JWTMiddleware) *gmux.Router {
	router := gmux.NewRouter().StrictSlash(false)
	apiRouter := router.PathPrefix("/api").Subrouter()

	initUnauthRoutes(apiRouter, webservice)
	initAuthRoutes(apiRouter, webservice, jwt)

	return router
}
