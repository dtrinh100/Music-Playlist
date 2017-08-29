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
func Login(rw http.ResponseWriter, req *http.Request) {
	var user model.User

	dec_err := json.NewDecoder(req.Body).Decode(&user)
	if common.HandleErrorWithMap(rw, dec_err, common.ErrMap{
		"Internal Server Error": "Failed to decode JSON",
	}, http.StatusInternalServerError) {
		return
	}

	user.Password = ""
	common.JsonStdResponse(user, rw)
}
