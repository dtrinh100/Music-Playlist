package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"github.com/dtrinh100/Music-Playlist/src/api/model"
	"golang.org/x/crypto/bcrypt"
	"github.com/dtrinh100/Music-Playlist/src/api/common"
)

// PostUser creates the user account
func PostUser(w http.ResponseWriter, req *http.Request, env *common.Env) error {

	decoder := json.NewDecoder(req.Body) // reads in request body
	var user model.User
	decErr := decoder.Decode(&user) // puts the JSON data into the user structure defined in the model package
	if decErr != nil {
		return StatusError{500, decErr}
	}
	if len(user.Username) < 2 || len(user.Username) > 30 {
		return StatusError{400, errors.New("usernames need to be more than 2 characters and less than 30 characters")}
	}
	emailRe := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`) // Regular expression to check for valid email, this is more strict than the Angular built-in validation
	if !emailRe.MatchString(user.Email) {
		return StatusError{400, errors.New("invalid email address")}
	}
	if len(user.Password) < 8 {
		return StatusError{400, errors.New("passwords need to be more at least 8 characters")}
	}

	// encrypts the password of the user before storing it in the db
	// recommended use a cost of 12 or more
	hashedPassword, genErr := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if genErr != nil {
		return StatusError{500, genErr}
	}

	insertErr := env.DB.InsertUser(user.Username, hashedPassword, user.Email)
	if insertErr != nil {
		return StatusError{500, insertErr}
	}

	userData, unmErr := json.Marshal(user)
	if unmErr != nil {
		return StatusError{500, unmErr}
	}

	defer req.Body.Close()

	w.Write(userData)

	return nil
}
