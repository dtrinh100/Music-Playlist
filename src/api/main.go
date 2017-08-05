package main

import (
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/alice"
	"github.com/dtrinh100/Music-Playlist/src/api/handler"
	"github.com/gorilla/mux"
)

func main() {

	// using Gorilla mux router instead of default one because it offers more flexibity
	r := mux.NewRouter()

	session, err := mgo.Dial("server1.example.com,server2.example.com")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	env := &handler.Env{
		DB: session,
	}

	// place middleware codes here, things like auth
	commonHandlers := alice.New()

	// adds the api prefix to all subroutes
	s := r.PathPrefix("/api/").Subrouter()

	// route handlers
	// ignore vet errors for unkeyed fields
	s.Handle("/users", commonHandlers.Then(handler.Handler{env, handler.PostUser})).Methods("POST")

	// start server on port 3000
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}
}
