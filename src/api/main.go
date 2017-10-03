package main

import (
	"log"
	"net/http"

	"github.com/dtrinh100/Music-Playlist/src/api/common"
	"github.com/dtrinh100/Music-Playlist/src/api/db"
	"github.com/dtrinh100/Music-Playlist/src/api/middleware"
	"github.com/dtrinh100/Music-Playlist/src/api/router"
)

/**
initAndGetHandler initializes the JWT key pair, the DB, the Env,
the route paths, & returns a handler.
*/
func initAndGetHandler() http.Handler {
	pub, priv := middleware.InitRSAKeyPair()

	dbConf := db.InitDB()
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
	serverConf := common.InitServer()

	server := &http.Server{
		Addr:    serverConf.Address,
		Handler: initAndGetHandler(),
	}

	if serverErr := server.ListenAndServe(); serverErr != nil {
		log.Fatal("Server failed to start:", serverErr)
	}
}
