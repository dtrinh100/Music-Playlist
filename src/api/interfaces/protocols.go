package interfaces

import "github.com/dtrinh100/Music-Playlist/src/api/usecases"

type DBHandler interface {
  Create(docs ...interface{}) error
  Update(selector usecases.M, update usecases.M) error
  Delete(selector usecases.M) error
  One(query usecases.M, result interface{}) error
  All(results interface{}) error
}

type UserInteractor interface {
  UpdatePasswordForEmail(userEmail, password string) error
  UpdateFirstNameForEmail(userEmail, firstName string) error
  UpdateLastNameForEmail(userEmail, lastName string) error
  UpdatePicURLForEmail(userEmail, picURL string) error
  // TODO: change contributions type to an array of songs?
  UpdateContributionsForEmail(userEmail string, contributions []interface{}) error
  // TODO: change playlist type to an array of songs?
  UpdatePlaylistForEmail(userEmail string, playlist []interface{}) error
  CreateNew(user *usecases.User) error
  RemoveByEmail(userEmail string) error
  GetByEmail(userEmail string) (*usecases.User, error)
  GetAll() ([]usecases.User, error)
}

type WebResponder interface {
	Success(http.ResponseWriter, usecases.M)
	Created(http.ResponseWriter, usecases.M)
	NoContent(http.ResponseWriter)

	Redirection(http.ResponseWriter)

	BadRequest(http.ResponseWriter, usecases.MPError)
	Unauthorized(http.ResponseWriter)
	Forbidden(http.ResponseWriter)
	NotFound(http.ResponseWriter)
	Gone(http.ResponseWriter)

	InternalServerError(http.ResponseWriter)
	ServiceUnavailable(http.ResponseWriter)
}
