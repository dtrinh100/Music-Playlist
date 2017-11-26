package interfaces

import (
  "net/http"
  "github.com/dtrinh100/Music-Playlist/src/api/usecases"
)

type WebserviceHandler struct {
  Logger         usecases.Logger
	UserInteractor UserInteractor
	Responder  WebResponder
}

func (webhandler WebserviceHandler) RegisterUser(rw http.ResponseWriter, req *http.Request) {
  // TODO: update to use JSON
  newUser := usecases.User{
    Username: req.FormValue("username"),
    Password: req.FormValue("password"),
    Email:    req.FormValue("email"),
  }

  if createErr := webhandler.UserInteractor.CreateNew(&newUser); createErr != nil {
    // TODO: qualify createErr and return appropriate response
    rw.Write([]byte("User Already Exists"))
    return
  }

  rw.Write([]byte("Registered successfully"))
}
