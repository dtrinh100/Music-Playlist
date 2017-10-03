package handler

import (
	"encoding/json"
	"time"

	"github.com/dtrinh100/Music-Playlist/src/api/common"
	"github.com/dtrinh100/Music-Playlist/src/api/middleware"
	"net/http"
)

func Logout(rw http.ResponseWriter, req *http.Request, env *common.Env) error {
	// TODO:	Add JWT to blacklist
	// Invalidate cookie
	middleware.SetSecuredCookie(rw, "none", time.Now())

	var res interface{}
	if unmarshalErr := json.Unmarshal([]byte(`{"Logout": "Successful"}`), &res); unmarshalErr != nil {
		return unmarshalErr
	}

	common.JSONStdResponse(rw, res)

	return nil
}
