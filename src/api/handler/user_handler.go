package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"github.com/dtrinh100/Music-Playlist/src/api/common"
	"github.com/dtrinh100/Music-Playlist/src/api/model"
	"golang.org/x/crypto/bcrypt"
)

// PostUser creates the user account
func PostUser(w http.ResponseWriter, req *http.Request, env *common.Env) error {
	decoder := json.NewDecoder(req.Body) // reads in request body
	var user model.User
	decErr := decoder.Decode(&user) // puts the JSON data into the user structure defined in the model package
	if decErr != nil {
		return decErr
	}
	if len(user.Username) < 2 || len(user.Username) > 30 {
		unKey := "Username"
		errMap := common.ErrMap{unKey: "Username must be greater than 3 and less than 30 characters"}
		return JSONError{
			Code:   http.StatusBadRequest,
			Err:    errors.New(errMap[unKey]),
			ErrMap: errMap,
		}
	}
	emailRe := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`) // Regular expression to check for valid email, this is more strict than the Angular built-in validation
	if !emailRe.MatchString(user.Email) {
		emailKey := "Email"
		errMap := common.ErrMap{emailKey: "Email format is invalid"}
		return JSONError{
			Code:   http.StatusBadRequest,
			Err:    errors.New(errMap[emailKey]),
			ErrMap: errMap,
		}
	}
	if len(user.Password) < 8 {
		pwKey := "Password"
		errMap := common.ErrMap{pwKey: "Password must be greater than 7 characters"}
		return JSONError{
			Code:   http.StatusBadRequest,
			Err:    errors.New(errMap[pwKey]),
			ErrMap: errMap,
		}
	}

	// encrypts the password of the user before storing it in the db
	// recommended use a cost of 12 or more
	hashedPassword, genErr := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if genErr != nil {
		return genErr
	}

	insertErr := env.DB.InsertUser(user.Username, hashedPassword, user.Email)
	if insertErr != nil {
		return genErr
	}

	userData, unmErr := json.Marshal(user)
	if unmErr != nil {
		return genErr
	}

	defer req.Body.Close()

	w.Write(userData)

	return nil
}
