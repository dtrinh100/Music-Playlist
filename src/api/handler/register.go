package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"github.com/dtrinh100/Music-Playlist/src/api/common"
	"github.com/dtrinh100/Music-Playlist/src/api/db"
	"github.com/dtrinh100/Music-Playlist/src/api/middleware"
	"github.com/dtrinh100/Music-Playlist/src/api/model"
)

// Register creates a user account by inserting it into the DB
func Register(rw http.ResponseWriter, req *http.Request, env *common.Env) error {
	var user model.User
	// Puts the JSON data into the user structure defined in the model package
	if decodeErr := json.NewDecoder(req.Body).Decode(&user); decodeErr != nil {
		return decodeErr
	}
	// Validates user's info
	if validationErr := validateUserInfo(&user); validationErr != nil {
		return validationErr
	}
	// Insert user into DB
	session := env.DB.GetSessionCopy()
	defer session.Close()

	ur := &db.UserRepository{DBName: env.DB.Name, Name: env.DB.UserTable.Name, Session: session}
	if regErr := ur.Register(&user); regErr != nil {
		errMap := common.ErrMap{"Register": "Username already exists"}
		return JSONError{
			Code:    http.StatusBadRequest,
			Err:     regErr,
			DataMap: common.ErrorList{Errors: errMap},
		}
	}
	// Creates & returns a valid JWT + user info
	jwtString, jwtExpireTime, jwtErr := middleware.GetJWT(env.RSAKeys.Private, user.Email)
	if jwtErr != nil {
		return jwtErr
	}

	middleware.SetSecuredCookie(rw, jwtString, jwtExpireTime)
	common.JSONStdResponse(rw, user)

	return nil
}

func validateUserInfo(user *model.User) error {
	// Validate username's length
	if len(user.Username) < 2 || len(user.Username) > 30 {
		unKey := "Username"
		errMap := common.ErrMap{unKey: "Username must be greater than 3 and less than 30 characters"}
		return JSONError{
			Code:    http.StatusBadRequest,
			Err:     errors.New(errMap[unKey]),
			DataMap: common.ErrorList{Errors: errMap},
		}
	}
	// Regular expression to check for valid email, this is more strict than the
	// Angular built-in validation
	emailRe := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRe.MatchString(user.Email) {
		emailKey := "Email"
		errMap := common.ErrMap{emailKey: "Email format is invalid"}
		return JSONError{
			Code:    http.StatusBadRequest,
			Err:     errors.New(errMap[emailKey]),
			DataMap: common.ErrorList{Errors: errMap},
		}
	}
	// Validate password's length
	if len(user.Password) < 8 {
		pwKey := "Password"
		errMap := common.ErrMap{pwKey: "Password must be greater than 7 characters"}
		return JSONError{
			Code:    http.StatusBadRequest,
			Err:     errors.New(errMap[pwKey]),
			DataMap: common.ErrorList{Errors: errMap},
		}
	}

	return nil
}
