package main

import (
	"log"
	"net/http"

	"github.com/dtrinh100/Music-Playlist/src/api/db"
	"github.com/dtrinh100/Music-Playlist/src/api/common"
	"github.com/dtrinh100/Music-Playlist/src/api/router"

	"gopkg.in/mgo.v2"
)

func main() {
	// Note: MPDatabase name comes from docker-compose.yml
	session, err := mgo.Dial("MPDatabase")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	dbc := db.InitDB(session)
	mainHandler := router.InitializeRoutes(eh)
	sc := common.InitServer()
	env := &common.Env{DB: dbc}
	mainHandler := router.InitializeRoutes(env)

	server := &http.Server{
		Addr:    sc.Address,
		Handler: mainHandler,
	}

	if server_err := server.ListenAndServe(); server_err != nil {
		log.Fatal("Server failed to start:", server_err)
	}
}
