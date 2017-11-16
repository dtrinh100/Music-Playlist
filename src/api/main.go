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
	dbUserHandler := infrastructure.NewMongoHandler(session, "dbName", "usertable")

	dbSongHandler := infrastructure.NewMongoHandler(session, "dbName", "songtable")

	handlers := make(map[string]interfaces.DBHandler)
	handlers["DBUserRepo"] = dbUserHandler
	handlers["DBSongRepo"] = dbSongHandler

	logger := new(infrastructure.Logger)

	userInteractor := new(usecases.UserInteractor)
	userInteractor.UserRepository = interfaces.NewDBUserRepo(handlers)
	userInteractor.Logger = logger

	webserviceHandler := new(interfaces.WebserviceHandler)
	webserviceHandler.UserInteractor = userInteractor
	webserviceHandler.Logger = logger

	// MIDDLEWARE

	jwt := new(middleware.JWTHandler)
	jwt.Logger = logger
	weblogger := new(middleware.WebLoggerHandler)
	weblogger.Logger = logger

	globalMiddlewares := []alice.Constructor{
		weblogger.Handle,
		jwt.Handle, // TODO: move this middleware to auth'ed routes later
	}

	// SERVER

	router := gmux.NewRouter().StrictSlash(false)
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/register", func(rw http.ResponseWriter, req *http.Request) {
		webserviceHandler.RegisterUser(rw, req)
	}).Methods("POST")

	server := &http.Server{
		Addr:    os.Getenv("MP_SRVRADDR_ENV"),
		Handler: alice.New(globalMiddlewares...).Then(router),
	}

	if serverErr := server.ListenAndServe(); serverErr != nil {
		logger.Log("Server failed to boot: " + serverErr.Error())
	}
}
