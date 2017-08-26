package main

import (
	"log"
	"net/http"

	"github.com/dtrinh100/Music-Playlist/src/api/db"
	"github.com/dtrinh100/Music-Playlist/src/api/handler"

	gmux "github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"gopkg.in/mgo.v2"
)

func main() {

	// using Gorilla mux router instead of default one because it offers more flexibity
	r := gmux.NewRouter()

	session, err := mgo.Dial("MPDatabase")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	env := &handler.Env{
		DB: &db.DB{Session: session},
	}

	// place middleware codes here, things like auth
	// TODO: replace .Classic() with .New() later.
	//		 For now, leave as .Classic() to benefit from logging.
	commonHandlers := negroni.Classic()
	commonHandlers.UseHandler(r)

	// adds the api prefix to all subroutes
	s := r.PathPrefix("/api/").Subrouter()

	// route handlers
	// ignore vet errors for unkeyed fields
	s.Handle("/users", handler.Handler{env, handler.PostUser}).Methods("POST")
	
	server := &http.Server{
		Addr: ":3000",
		Handler: commonHandlers,
	}

	if server_err := server.ListenAndServe(); server_err {
		log.Fatal("Server failed to start:", server_err)
	}
}
