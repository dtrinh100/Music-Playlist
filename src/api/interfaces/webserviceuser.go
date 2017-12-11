package interfaces

import (
  "net/http"
  "github.com/dtrinh100/Music-Playlist/src/api/usecases"
  gmux "github.com/gorilla/mux"
  "encoding/json"
  "io"
)

func (webhandler *WebserviceHandler) Users(rw http.ResponseWriter, req *http.Request) {
  users, usersErr := webhandler.UserInteractor.GetAll()

  if usersErr != nil {
    webhandler.Responder.InternalServerError(rw)
    return
  }

  webhandler.Responder.Success(rw, usecases.M{"users": users})
}

func (webhandler *WebserviceHandler) User(rw http.ResponseWriter, req *http.Request) {
  username := gmux.Vars(req)["username"]

  user, userErr := webhandler.UserInteractor.GetByUsername(username)

  if userErr != nil {
    switch userErr.Status() {
    case usecases.UserFaultErr:
      webhandler.Responder.BadRequest(rw, userErr)
    default:
      webhandler.Responder.InternalServerError(rw)
    }
    return
  }

  webhandler.Responder.Success(rw, usecases.M{"user": user})
}

func (webhandler *WebserviceHandler) UpdateUser(rw http.ResponseWriter, req *http.Request) {
  username := gmux.Vars(req)["username"]
  authedEmail := req.Context().Value("mpEmailKey").(string)

  if webhandler.isUnauthorized(rw, username, authedEmail) {
    return
  }

  updateDict := usecases.M{}

  if decErr := json.NewDecoder(req.Body).Decode(&updateDict); decErr != nil && decErr != io.EOF {
    webhandler.Responder.InternalServerError(rw)
    return
  }

  badRequestFn := func(msg string) {
    webhandler.Responder.BadRequest(rw, &usecases.FaultError{
      FaultEntity: usecases.UserFaultErr,
      Message:     msg})
  }

  if len(updateDict) == 0 {
    badRequestFn("empty updates")
    return
  }

  for key, value := range updateDict {
    val := value.(string)
    switch key {
    case "firstname":
      webhandler.UserInteractor.UpdateFirstNameForEmail(authedEmail, val)
    case "lastname":
      webhandler.UserInteractor.UpdateLastNameForEmail(authedEmail, val)
    case "password":
      webhandler.UserInteractor.UpdatePasswordForEmail(authedEmail, val)
    case "picurl":
      webhandler.UserInteractor.UpdatePicURLForEmail(authedEmail, val)
      // TODO: handle contributions & playlist
    default:
      badRequestFn("cannot update - " + key)
      return
    }
  }

  webhandler.Responder.NoContent(rw)
}

func (webhandler *WebserviceHandler) DeleteUser(rw http.ResponseWriter, req *http.Request) {
  username := gmux.Vars(req)["username"]
  authedEmail := req.Context().Value("mpEmailKey").(string)

  if webhandler.isUnauthorized(rw, username, authedEmail) {
    return
  }

  if removeErr := webhandler.UserInteractor.RemoveByEmail(authedEmail); removeErr != nil {
    webhandler.Responder.NoContent(rw)
    return
  }

  webhandler.Responder.NoContent(rw)
}

func (webhandler *WebserviceHandler) isUnauthorized(rw http.ResponseWriter, username, authedEmail string) bool {
  requestedUser, userErr := webhandler.UserInteractor.GetByUsername(username)

  if userErr != nil {
    switch userErr.Status() {
    case usecases.UserFaultErr:
      webhandler.Responder.BadRequest(rw, userErr)
    default:
      webhandler.Responder.InternalServerError(rw)
    }
    return true
  }

  if authedEmail != requestedUser.Email {
    webhandler.Responder.Unauthorized(rw)
    return true
  }

  return false
}