package handler

import (
	"net/http"
	"github.com/dtrinh100/Music-Playlist/src/api/model"
	"encoding/json"
	"github.com/dtrinh100/Music-Playlist/src/api/common"
)

func Login(rw http.ResponseWriter, req *http.Request) {
	var user model.User

	dec_err := json.NewDecoder(req.Body).Decode(&user)
	if common.HandleErrorWithMessage(rw, dec_err, common.ErrMap{
		"Internal Server Error": "Failed to decode JSON",
	}, http.StatusInternalServerError) {
		return
	}

	user.Password = ""
	common.JsonStdResponse(user, rw)
}
