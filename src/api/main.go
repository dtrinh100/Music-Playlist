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

	dbName := "musicplaylistdb"
	dbUserHandler := infrastructure.NewMongoHandler(session, dbName, "usertable")
	dbSongHandler := infrastructure.NewMongoHandler(session, dbName, "songtable")
	dbCounterHandler := infrastructure.NewMongoHandler(session, dbName, "countertable")

	songSeq := interfaces.Counter{"songid", 0}
	userSeq := interfaces.Counter{"userid", 0}

	dbCounterHandler.Create(songSeq)
	dbCounterHandler.Create(userSeq)

	handlers := make(map[string]interfaces.DBHandler)
	handlers["DBUserRepo"] = dbUserHandler
	handlers["DBSongRepo"] = dbSongHandler
	handlers["DBCounterRepo"] = dbCounterHandler

	logger := new(infrastructure.Logger)

	userInteractor := new(usecases.UserInteractor)
	userInteractor.UserRepository = interfaces.NewDBUserRepo(handlers)
	userInteractor.Logger = logger

	jsonResponder := new(infrastructure.JSONWebResponder)
	songInteractor := new(usecases.SongInteractor)
	songInteractor.SongRepository = interfaces.NewDBSongRepo(handlers)
	songInteractor.Logger = logger

	jwtHandler := interfaces.NewJWTHandler()

	webserviceHandler := &interfaces.WebserviceHandler{
		SongInteractor: songInteractor,
		UserInteractor: userInteractor,
		Responder:      jsonResponder,
		JWTHandler:     jwtHandler,
	}

	// MIDDLEWARE

	weblogger := new(middleware.WebLoggerMiddleware)
	weblogger.Logger = logger
	weblogger.Responder = jsonResponder
	globalMiddlewares := []alice.Constructor{weblogger.Handle}

	jwt := new(middleware.JWTMiddleware)
	jwt.Logger = logger
	jwt.Responder = jsonResponder
	jwt.JWTHandler = jwtHandler

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
	authRouter := gmux.NewRouter()
	authHandle := alice.New(jwt.Handle).Then(authRouter)
	apiRouter := router.PathPrefix("/api").Subrouter()
	// Note: the following 2 lines are needed to allow sub-routers to be used as a handler
	apiRouter.Handle("/auth", authHandle)
	apiRouter.Handle("/auth/{path:.*}", authHandle)

	// TODO: figure out if there's a RESTful way of doing authentication
	apiRouter.HandleFunc("/register", func(rw http.ResponseWriter, req *http.Request) {
		webservice.RegisterUser(rw, req)
	}).Methods("POST")

	apiRouter.HandleFunc("/login", func(rw http.ResponseWriter, req *http.Request) {
		webservice.LoginUser(rw, req)
	}).Methods("POST")

	// =-=-=-=-=-=-=-=-=-= USER ROUTER

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

	// =-=-=-=-=-=-=-=-=-= SONG ROUTER

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
