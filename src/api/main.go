package main

import (
	"github.com/dtrinh100/Music-Playlist/src/api/infrastructure/middleware"
	"github.com/dtrinh100/Music-Playlist/src/api/infrastructure"
	"github.com/dtrinh100/Music-Playlist/src/api/interfaces"
	"net/http"
	"github.com/justinas/alice"
	"os"
)

func main() {
	// WEBSERVICE

	handlers := infrastructure.GetAndInitDBHandlersForDBName(os.Getenv("MP_DBNAME_ENV"))
	logger := new(infrastructure.Logger)
	userInteractor := infrastructure.GetAndInitUserInteractor(logger, handlers)
	songInteractor := infrastructure.GetAndInitSongInteractor(logger, handlers)
	jwtHandler := interfaces.NewJWTHandler()
	jsonResponder := new(infrastructure.JSONWebResponder)

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

	router := infrastructure.GetRouterWithRoutes(webserviceHandler, jwt)

	server := &http.Server{
		Addr:    os.Getenv("MP_SRVRADDR_ENV"),
		Handler: alice.New(globalMiddlewares...).Then(router),
	}

	if serverErr := server.ListenAndServe(); serverErr != nil {
		logger.Log("Server failed to boot: " + serverErr.Error())
	}
}


