package interfaces

import (
	"net/http"
	"github.com/dtrinh100/Music-Playlist/src/api/usecases"
	"encoding/json"
	"regexp"
	"github.com/dtrinh100/Music-Playlist/src/api/domain"
)

type WebserviceHandler struct {
	SongInteractor SongInteractor
	UserInteractor UserInteractor

	Responder  WebResponder
	JWTHandler JWTHandler
}

func (webhandler *WebserviceHandler) RegisterUser(rw http.ResponseWriter, req *http.Request) {
	validateUserInfoFn := func(user *usecases.User) usecases.MPError {
		respFn := func(msg string) usecases.MPError {
			return &usecases.FaultError{usecases.UserFaultErr, msg}
		}

		// Validate username's length
		if len(user.Username) < 2 || len(user.Username) > 30 {
			return respFn("Username must be greater than 3 and less than 30 characters")
		}
		// Regular expression to check for valid email, this is more strict than the
		// Angular built-in validation
		if !regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(user.Email) {
			return respFn("Email format is invalid")
		}
		// Validate password's length
		if len(user.Password) < 8 {
			return respFn("Password must be greater than 7 characters")
		}

		return nil
	}
	var newUser usecases.User

	if decodeErr := json.NewDecoder(req.Body).Decode(&newUser); decodeErr != nil {
		webhandler.Responder.InternalServerError(rw)
		return
	}

	if validateErr := validateUserInfoFn(&newUser); validateErr != nil {
		webhandler.Responder.BadRequest(rw, validateErr)
		return
	}

	if createErr := webhandler.UserInteractor.CreateNew(&newUser); createErr != nil {
		switch createErr.Status() {
		case usecases.UserFaultErr:
			webhandler.Responder.BadRequest(rw, createErr)
		default:
			webhandler.Responder.InternalServerError(rw)
		}
		return
	}

	newUser.Contributions = []domain.Song{}
	newUser.Playlist = []domain.Song{}

	webhandler.JWTHandler.ValidateUserEmail(rw, newUser.Email)
	webhandler.Responder.Success(rw, usecases.M{"user": newUser})
}

func (webhandler *WebserviceHandler) LoginUser(rw http.ResponseWriter, req *http.Request) {
	var possibleUser usecases.User

	if decodeErr := json.NewDecoder(req.Body).Decode(&possibleUser); decodeErr != nil {
		webhandler.Responder.InternalServerError(rw)
		return
	}

	existingUser, getErr := webhandler.UserInteractor.GetByEmail(possibleUser.Email)

	if getErr != nil {
		webhandler.Responder.Unauthorized(rw)
		return
	}

	compareErr := webhandler.UserInteractor.ComparePassword(existingUser, possibleUser.Password)
	if compareErr != nil {
		webhandler.Responder.Unauthorized(rw)
		return
	}

	webhandler.JWTHandler.ValidateUserEmail(rw, existingUser.Email)
	webhandler.Responder.Success(rw, usecases.M{"user": existingUser})
}

func (webhandler *WebserviceHandler) VerifyUser(rw http.ResponseWriter, req *http.Request) {
	authedEmail := req.Context().Value("mpEmailKey").(string)

	user, userErr := webhandler.UserInteractor.GetByEmail(authedEmail)

	if userErr != nil {
		webhandler.Responder.Unauthorized(rw)
	}

	webhandler.Responder.Success(rw, usecases.M{"user": user})
}

func (webhandler *WebserviceHandler) LogoutUser(rw http.ResponseWriter, req *http.Request) {
	// TODO: blacklist JWT
	webhandler.JWTHandler.InvalidateUserEmail(rw)
	webhandler.Responder.NoContent(rw)
}
