package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func main() {

	// using Gorilla mux router instead of default one because it offers more flexibity
	r := mux.NewRouter()

	// place middleware codes here
	commonHandlers := alice.New()

	// adds the api prefix to all subroutes
	s := r.PathPrefix("/api/").Subrouter()

	// route handlers
	s.Handle("/users", commonHandlers.ThenFunc()).Methods("POST")

	// start server on port 3000
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}
}
