package main

import (
	"log"
	"net/http"

	"github.com/dtrinh100/Music-Playlist/src/api/db"
	"github.com/dtrinh100/Music-Playlist/src/api/handler"
	"github.com/dtrinh100/Music-Playlist/src/api/common"
	"github.com/dtrinh100/Music-Playlist/src/api/router"

	"gopkg.in/mgo.v2"
)

func init() {
	common.InitServer()
}

func main() {
	// TODO: initialize handler.Env & DB in common.InitServer() later
	session, err := mgo.Dial("MPDatabase")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	env := &handler.Env{
		DB: &db.DB{Session: session},
	}

	mainHandler := router.InitializeRoutes(env)

	server := &http.Server{
		Addr:    common.AppConfig.Server.Address,
		Handler: mainHandler,
	}

	if server_err := server.ListenAndServe(); server_err != nil {
		log.Fatal("Server failed to start:", server_err)
	}
}
