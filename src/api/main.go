package main

import (
	"log"
	"net/http"

	"github.com/dtrinh100/Music-Playlist/src/api/common"
	"github.com/dtrinh100/Music-Playlist/src/api/db"
	"github.com/dtrinh100/Music-Playlist/src/api/middleware"
	"github.com/dtrinh100/Music-Playlist/src/api/router"

	"gopkg.in/mgo.v2"
)

// Note: 'MPDatabase' name comes from docker-compose.yml
const dbURLAddress = "MPDatabase"

/**
initAndGetHandler initializes the JWT key pair, the DB, the Env,
the route paths, & returns a handler.
*/
func initAndGetHandler(session *mgo.Session) http.Handler {
	pub, priv := middleware.InitRSAKeyPair()

	dbConf := db.InitDB(session)
	env := &common.Env{
		DB: dbConf,
		RSAKeys: common.RSAKeys{
			Public:  pub,
			Private: priv,
		},
	}

	return router.InitializeRoutes(env)
}

/**
main is the entry-function of the API.
*/
func main() {
	session, sessionErr := mgo.Dial(dbURLAddress)
	common.Fatal(sessionErr, "Failed to obtain a DB session")
	defer session.Close()

	serverConf := common.InitServer()

	server := &http.Server{
		Addr:    serverConf.Address,
		Handler: initAndGetHandler(session),
	}

	if serverErr := server.ListenAndServe(); serverErr != nil {
		log.Fatal("Server failed to start:", serverErr)
	}
}
