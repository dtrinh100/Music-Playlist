package handler

import (
	"net/http"
	"github.com/dtrinh100/Music-Playlist/src/api/model"
	"encoding/json"
	"github.com/dtrinh100/Music-Playlist/src/api/common"
)

/**
	TODO: update to use the DB.
	This function logs-in a user with valid credentials.

	Path: [POST] '/api/auth'
*/
func Login(rw http.ResponseWriter, req *http.Request, env *Env) error {
	var user model.User

	dec_err := json.NewDecoder(req.Body).Decode(&user)
	if dec_err != nil {
		return dec_err
	}

	user.Password = ""
	common.JsonStdResponse(rw, user)

	return nil
}
