package handler

import (

	"github.com/dtrinh100/Music-Playlist/src/api/common"
	"github.com/dtrinh100/Music-Playlist/src/api/db"
	"github.com/dtrinh100/Music-Playlist/src/api/middleware"
	"github.com/dtrinh100/Music-Playlist/src/api/model"

	"net/http"

	"encoding/json"
)

/**
Login authenticates a user with valid credentials.

Path: [POST] '/api/auth/login'
*/
func Login(rw http.ResponseWriter, req *http.Request, env *common.Env) error {
	var user *model.User
	// Puts the JSON data into the user structure defined in the model package
	if decodeErr := json.NewDecoder(req.Body).Decode(user); decodeErr != nil {
		return decodeErr
	}

	session := env.DB.GetSessionCopy()
	defer session.Close()

	ur := db.UserRepository{DBName: env.DB.Name, Name: env.DB.UserTable.Name, Session: session}
	if loginErr := ur.Login(user); loginErr != nil {
		return JSONError{
			Code: http.StatusUnauthorized,
			Err:  loginErr,
			DataMap: common.ErrorList{
				Errors: common.ErrMap{"Login": "Invalid Credentials"},
			},
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
