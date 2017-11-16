package usecases

import "github.com/dtrinh100/Music-Playlist/src/api/domain"

type User struct {
	Username       string        `json:"username" bson:"username"`
	Password       string        `json:"password,omitempty" bson:"-"`
	HashedPassword []byte        `json:"-" bson:"hashedpassword"`
	Email          string        `json:"email" bson:"email"`
	FirstName      string        `json:"firstname" bson:"firstname"`
	LastName       string        `json:"lastname" bson:"lastname"`
	PicURL         string        `json:"picurl" bson:"picurl"`
	Contributions  []domain.Song `json:"contributions" bson:"contributions"`
	Playlist       []domain.Song `json:"playlist" bson:"playlist"`
}
