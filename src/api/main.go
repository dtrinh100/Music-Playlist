package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func main() {

	// using Gorilla mux router instead of default one because it offers more flexibity
	mx := mux.NewRouter()

	// place middleware codes here
	commonHandlers := alice.New()

	// route handlers
	mx.Handle("/api/register", commonHandlers.ThenFunc())

	// start server on port 3000
	err := http.ListenAndServe(":3000", mx)
	if err != nil {
		log.Fatal(err)
	}
}
