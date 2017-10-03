package handler

import (
	"github.com/dtrinh100/Music-Playlist/src/api/common"
	"github.com/dtrinh100/Music-Playlist/src/api/db"
	"github.com/dtrinh100/Music-Playlist/src/api/middleware"
	"github.com/dtrinh100/Music-Playlist/src/api/model"
	"net/http"
)

/**
Verify authenticates a user through the JWT in the cookie.

Path: [POST] '/api/auth/verify'
*/
func Verify(rw http.ResponseWriter, req *http.Request, env *common.Env) error {
	claims := req.Context().Value(middleware.MPClaimsKey).(model.AppClaims)
	session := env.DB.GetSessionCopy()
	defer session.Close()

	ur := db.UserRepository{DBName: env.DB.Name, Name: env.DB.UserTable.Name, Session: session}
	user2Verify := &model.User{Email: claims.UserEmail}

	if verifyErr := ur.Verify(user2Verify); verifyErr != nil {
		errMap := common.ErrMap{"Credentials": "Invalid Or Expired Credentials"}
		return JSONError{
			Code:    http.StatusUnauthorized,
			Err:     verifyErr,
			DataMap: errMap,
		}
	}

	common.JSONStdResponse(rw, user2Verify)

	return nil
}
