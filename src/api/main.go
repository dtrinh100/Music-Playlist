package main

import (
	"github.com/dtrinh100/Music-Playlist/src/api/infrastructure/middleware"
	"github.com/dtrinh100/Music-Playlist/src/api/infrastructure"
	"github.com/dtrinh100/Music-Playlist/src/api/interfaces"
	"github.com/dtrinh100/Music-Playlist/src/api/usecases"
	gmux "github.com/gorilla/mux"
	"net/http"
	"github.com/justinas/alice"
	"os"
)

func main() {
	// WEBSERVICE

	session := infrastructure.NewMongoSession("MPDatabase", "", "")
	dbUserHandler := infrastructure.NewMongoHandler(session, "musicplaylistdb", "usertable")
	dbSongHandler := infrastructure.NewMongoHandler(session, "musicplaylistdb", "songtable")

	handlers := make(map[string]interfaces.DBHandler)
	handlers["DBUserRepo"] = dbUserHandler
	handlers["DBSongRepo"] = dbSongHandler

	logger := new(infrastructure.Logger)

	userInteractor := new(usecases.UserInteractor)
	userInteractor.UserRepository = interfaces.NewDBUserRepo(handlers)
	userInteractor.Logger = logger

	jsonResponder := new(infrastructure.JSONWebResponder)

	jwtHandler := interfaces.NewJWTHandler()

	webserviceHandler := &interfaces.WebserviceHandler{
		userInteractor,
		jsonResponder,
		jwtHandler,
	}

	// MIDDLEWARE

	jwt := new(middleware.JWTMiddleware)
	jwt.Logger = logger
	jwt.Responder = jsonResponder
	jwt.JWTHandler = jwtHandler

	weblogger := new(middleware.WebLoggerMiddleware)
	weblogger.Logger = logger
	weblogger.Responder = jsonResponder

	globalMiddlewares := []alice.Constructor{
		weblogger.Handle,
	}

	// SERVER

	router := gmux.NewRouter().StrictSlash(false)
	initRoutes(router, webserviceHandler, jwt)

	server := &http.Server{
		Addr:    os.Getenv("MP_SRVRADDR_ENV"),
		Handler: alice.New(globalMiddlewares...).Then(router),
	}

	if serverErr := server.ListenAndServe(); serverErr != nil {
		logger.Log("Server failed to boot: " + serverErr.Error())
	}
}

func initRoutes(router *gmux.Router, webservice *interfaces.WebserviceHandler, jwt *middleware.JWTMiddleware) {
	apiRouter := router.PathPrefix("/api").Subrouter()
	// TODO: figure out if there's a RESTful way of doing authentication
	apiRouter.HandleFunc("/register", func(rw http.ResponseWriter, req *http.Request) {
		webservice.RegisterUser(rw, req)
	}).Methods("POST")

	apiRouter.HandleFunc("/login", func(rw http.ResponseWriter, req *http.Request) {
		webservice.LoginUser(rw, req)
	}).Methods("POST")

	authRouter := apiRouter.PathPrefix("/auth").Subrouter()

	authRouteFn := func(path string, fn func(http.ResponseWriter, *http.Request)) *gmux.Route {
		return authRouter.Handle(path, alice.New(jwt.Handle).Then(http.HandlerFunc(fn)))
	}

	authRouteFn("/logout", func(rw http.ResponseWriter, req *http.Request) {
		webservice.LogoutUser(rw, req)
	}).Methods("POST")
}
