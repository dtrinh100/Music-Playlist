package interfaces

import (
	"github.com/dtrinh100/Music-Playlist/src/api/domain"
	"github.com/dtrinh100/Music-Playlist/src/api/usecases"
)

type DBSongRepo struct {
	DBRepo
}

func (repo *DBSongRepo) OneByID(songID int) (*domain.Song, error) {
	song := new(domain.Song)

	oneErr := repo.dbHandler.One(usecases.M{"_id": songID}, song)

	return song, oneErr
}

func (repo *DBSongRepo) All() ([]domain.Song, error) {
	songs := []domain.Song{}
	allErr := repo.dbHandler.All(&songs)

	return songs, allErr
}

func (repo *DBSongRepo) Create(song *domain.Song) error {
	songID, seqErr := repo.getNextSequence("songid")

	if seqErr != nil {
		return seqErr
	}

	song.ID = songID

	if createErr := repo.dbHandler.Create(*song); createErr != nil {
		return createErr
	}

	return nil
}

func (repo *DBSongRepo) Update(songID int, changes usecases.M) error {
	return repo.dbHandler.Update(usecases.M{"_id": songID}, changes)
}

func (repo *DBSongRepo) Delete(songID int) error {
	return repo.dbHandler.Delete(usecases.M{"_id": songID})
}

func NewDBSongRepo(dbHandlers map[string]DBHandler) *DBSongRepo {
	dbSongRepo := new(DBSongRepo)
	dbSongRepo.dbHandlers = dbHandlers
	dbSongRepo.dbHandler = dbHandlers["DBSongRepo"]

	return dbSongRepo
}
