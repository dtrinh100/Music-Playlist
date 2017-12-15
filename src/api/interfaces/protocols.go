package interfaces

import (
	"github.com/dtrinh100/Music-Playlist/src/api/usecases"
	"github.com/dtrinh100/Music-Playlist/src/api/domain"
	"net/http"
)

type DBHandler interface {
	EnsureIndex(fieldName string) error
	FindAndModify(query, update usecases.M, result interface{}) (interface{}, error)
	Create(docs ...interface{}) error
	Update(selector usecases.M, update usecases.M) error
	Delete(selector usecases.M) error
	One(query usecases.M, result interface{}) error
	All(results interface{}) error
}

type UserInteractor interface {
	UpdatePasswordForEmail(userEmail, password string) usecases.MPError
	UpdateFirstNameForEmail(userEmail, firstName string) usecases.MPError
	UpdateLastNameForEmail(userEmail, lastName string) usecases.MPError
	UpdatePicURLForEmail(userEmail, picURL string) usecases.MPError
	UpdateContributionsForEmail(userEmail string, contributions []domain.Song) usecases.MPError
	UpdatePlaylistForEmail(userEmail string, playlist []domain.Song) usecases.MPError
	
	CreateNew(user *usecases.User) usecases.MPError
	RemoveByEmail(userEmail string) usecases.MPError
	GetByEmail(userEmail string) (*usecases.User, usecases.MPError)
	GetByUsername(username string) (*usecases.User, usecases.MPError)
	GetAll() ([]usecases.User, usecases.MPError)

	ComparePassword(user *usecases.User, clearTextPass string) usecases.MPError
}

type SongInteractor interface {
	UpdateDislikes(songID int, dislikes int) usecases.MPError
	UpdateLikes(songID int, likes int) usecases.MPError
	UpdateName(songID int, songName string) usecases.MPError
	UpdateState(songID int, state string) usecases.MPError
	UpdateCountry(songID int, country string) usecases.MPError

	All() ([]domain.Song, usecases.MPError)
	GetByID(songID int) (*domain.Song, usecases.MPError)
	Create(song *domain.Song) usecases.MPError
	Delete(songID int) usecases.MPError
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
