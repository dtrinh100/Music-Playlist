package handler

import (
	"encoding/json"
	"github.com/dtrinh100/Music-Playlist/src/api/common"
	"github.com/dtrinh100/Music-Playlist/src/api/model"
	"net/http"
)

/**
TODO: update to use the DB.
Login authenticates a user with valid credentials.

Path: [POST] '/api/auth'
*/
func Login(rw http.ResponseWriter, req *http.Request, env *common.Env) error {
	var user model.User

	decErr := json.NewDecoder(req.Body).Decode(&user)
	if decErr != nil {
		return decErr
	}

	user.Password = ""

	common.JSONStdResponse(rw, user)
	return nil
}
