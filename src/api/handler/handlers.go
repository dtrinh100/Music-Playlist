package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dtrinh100/Music-Playlist/src/api/model"

	"gopkg.in/mgo.v2"
)

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}

// Move this into own file later on
type Env struct {
	DB *mgo.Session
}

// The Handler struct that takes a configured Env and a function matching
// our useful signature.
type Handler struct {
	*Env
	H func(e *Env, w http.ResponseWriter, r *http.Request) error
}

// ServeHTTP allows our Handler type to satisfy http.Handler.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(h.Env, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}

// PostUser creates the user account
func PostUser(env *Env, w http.ResponseWriter, req *http.Request) error {

	decoder := json.NewDecoder(req.Body) // reads in request body
	var user model.User
	err := decoder.Decode(&user) // puts the JSON data into the user structure defined in the model package

	userData, err := json.Marshal(user)

	defer req.Body.Close()
	if err != nil {
		log.Print(err) // In real life, we would store this data remotely
		//	w.Write()      // TODO: figure out how to check types seperately
		return StatusError{500, err}
	}

	w.Write(userData)

	return nil
}
