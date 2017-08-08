package main

import (
	"log"
	"net/http"

	"github.com/dtrinh100/Music-Playlist/src/api/handler"
	"github.com/justinas/alice"
	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

func main() {

	// using Gorilla mux router instead of default one because it offers more flexibity
	r := mux.NewRouter()

	//	session, err := mgo.Dial(os.Getenv("DB_CONTAINER_NAME"))
	//	if err != nil {
	//		panic(err)
	//	}
	//defer session.Close()

	env := &handler.Env{
		DB: nil, // TODO: REMEMBER TO REPLACE THIS WITH THE MONGODB SESSION
	}

	// place middleware codes here, things like auth
	commonHandlers := alice.New()

	// adds the api prefix to all subroutes
	s := r.PathPrefix("/api/").Subrouter()

	// route handlers
	// ignore vet errors for unkeyed fields
	s.Handle("/users", commonHandlers.Then(handler.Handler{env, handler.PostUser})).Methods("POST")
	// TODO: remove CORS after integration with Caddy
	handlers := cors.Default().Handler(r)

	// start server on port 3000
	err := http.ListenAndServe(":3000", handlers)
	if err != nil {
		log.Fatal(err)
	}
}
