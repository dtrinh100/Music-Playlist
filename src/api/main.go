package main

import (
	"log"
	"net/http"

	"github.com/dtrinh100/Music-Playlist/src/api/db"
	"github.com/dtrinh100/Music-Playlist/src/api/common"
	"github.com/dtrinh100/Music-Playlist/src/api/router"

	"gopkg.in/mgo.v2"
)

/**
	main is the entry-function of the api.
*/
func main() {
	// Note: MPDatabase name comes from docker-compose.yml
	session, dialErr := mgo.Dial("MPDatabase")
	if dialErr != nil {
		panic(dialErr)
	}
	defer session.Close()

	serverConf := common.InitServer()
	dbConf := db.InitDB(session)
	env := &common.Env{DB: dbConf}
	mainHandler := router.InitializeRoutes(env)

	server := &http.Server{
		Addr:    serverConf.Address,
		Handler: mainHandler,
	}

	if serverErr := server.ListenAndServe(); serverErr != nil {
		log.Fatal("Server failed to start:", serverErr)
	}
}
