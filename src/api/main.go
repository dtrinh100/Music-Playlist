package main

import (
	"log"
	"net/http"

	"github.com/dtrinh100/Music-Playlist/src/api/db"
	"github.com/dtrinh100/Music-Playlist/src/api/handler"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"gopkg.in/mgo.v2"
)

func main() {

	// using Gorilla mux router instead of default one because it offers more flexibity
	r := mux.NewRouter()

	session, err := mgo.Dial("MPDatabase")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	env := &handler.Env{
		DB: &db.DB{Session: session},
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
