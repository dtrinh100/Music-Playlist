package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dtrinh100/Music-Playlist/src/api/model"
	"golang.org/x/crypto/bcrypt"
)

// PostUser creates the user account
func PostUser(env *Env, w http.ResponseWriter, req *http.Request) error {

	decoder := json.NewDecoder(req.Body) // reads in request body
	var user model.User
	err := decoder.Decode(&user) // puts the JSON data into the user structure defined in the model package
	if err != nil {
		log.Print(err)
		return StatusError{500, err}
	}
	// encrypts the password of the user before storing it in the db
	// recommended use a cost of 12 or more
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		log.Print(err)
		return StatusError{500, err}
	}

	env.DB.InsertUser(user.Username, hashedPassword, user.Email)

	userData, err := json.Marshal(user)
	if err != nil {
		log.Print(err)
		return StatusError{500, err}
	}

	defer req.Body.Close()
	if err != nil {
		log.Print(err) // In real life, we would store this data remotely
		//	w.Write()      // TODO: figure out how to check types seperately
		return StatusError{500, err}
	}

	w.Write(userData)

	return nil
}
