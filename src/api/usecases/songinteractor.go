package usecases

import (
	"github.com/dtrinh100/Music-Playlist/src/api/domain"
)

type SongInteractor struct {
	SongRepository SongRepository
	Logger         Logger
}

// TODO: UserInteractor also uses something similar to this. DRY.
func (interactor *SongInteractor) errorHandler(err error) MPError {
	if err != nil {
		interactor.Logger.Log("Error occurred: " + err.Error())
		return &FaultError{AppFaultErr, err.Error()}
	}

	return nil
}

func (interactor *SongInteractor) All() ([]domain.Song, MPError) {
	songs, allErr := interactor.SongRepository.All()
	return songs, interactor.errorHandler(allErr)
}

func (interactor *SongInteractor) GetByID(songID int) (*domain.Song, MPError) {
	song, oneErr := interactor.SongRepository.OneByID(songID)

	if oneErr != nil && (oneErr.Error() == "not found" || oneErr.Error() == "invalid syntax") {
		interactor.Logger.Log("Song ID not found")
		return nil, &FaultError{UserFaultErr, oneErr.Error()}
	}

	return song, interactor.errorHandler(oneErr)
}

func (interactor *SongInteractor) Create(song *domain.Song) MPError {
	interactor.Logger.Log("Attempting to create song: " + song.Name)
	err := interactor.SongRepository.Create(song)
	return interactor.errorHandler(err)
}

func (interactor *SongInteractor) UpdateName(songID int, songName string) MPError {
	err := interactor.SongRepository.Update(songID, M{"name": songName})
	return interactor.errorHandler(err)
}

func (interactor *SongInteractor) UpdateState(songID int, state string) MPError {
	err := interactor.SongRepository.Update(songID, M{"state": state})
	return interactor.errorHandler(err)
}

func (interactor *SongInteractor) UpdateCountry(songID int, country string) MPError {
	err := interactor.SongRepository.Update(songID, M{"country": country})
	return interactor.errorHandler(err)
}

func (interactor *SongInteractor) UpdateLikes(songID int, likes int) MPError {
	err := interactor.SongRepository.Update(songID, M{"likes": likes})
	return interactor.errorHandler(err)
}

func (interactor *SongInteractor) UpdateDislikes(songID int, dislikes int) MPError {
	err := interactor.SongRepository.Update(songID, M{"dislikes": dislikes})
	return interactor.errorHandler(err)
}

func (interactor *SongInteractor) Delete(songID int) MPError {
	err := interactor.SongRepository.Delete(songID)
	return interactor.errorHandler(err)
}
