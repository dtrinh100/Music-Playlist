package usecases

import (
	"github.com/dtrinh100/Music-Playlist/src/api/domain"
)

type UserInteractor struct {
  UserRepository UserRepository
  Logger         Logger
}

func (interactor *UserInteractor) errorHandler(err error) MPError {
	if err != nil {
		interactor.Logger.Log("Error occurred: " + err.Error())
		return &FaultError{AppFaultErr, err.Error()}
	}

	return nil
}

func (interactor *UserInteractor) UpdatePasswordForEmail(userEmail, password string) MPError {
	err := interactor.UserRepository.Update(userEmail, M{"password": password})
	return interactor.errorHandler(err)
}

func (interactor *UserInteractor) UpdateFirstNameForEmail(userEmail, firstName string) MPError {
	err := interactor.UserRepository.Update(userEmail, M{"firstname": firstName})
	return interactor.errorHandler(err)
}

func (interactor *UserInteractor) UpdateLastNameForEmail(userEmail, lastName string) MPError {
	err := interactor.UserRepository.Update(userEmail, M{"lastname": lastName})
	return interactor.errorHandler(err)
}

func (interactor *UserInteractor) UpdatePicURLForEmail(userEmail, picURL string) MPError {
	err := interactor.UserRepository.Update(userEmail, M{"picurl": picURL})
	return interactor.errorHandler(err)
}

// TODO: implement & unit-test
func (interactor *UserInteractor) UpdateContributionsForEmail(userEmail string, contributions []domain.Song) MPError {
	return nil
}

// TODO: implement & unit-test
func (interactor *UserInteractor) UpdatePlaylistForEmail(userEmail string, playlist []domain.Song) MPError {
	return nil
}

func (interactor *UserInteractor) CreateNew(user *User) MPError {
	if existingUser, oneErr := interactor.GetByEmail(user.Email); oneErr == nil && existingUser != nil {
		msg := "user already exists"
		interactor.Logger.Log(msg + ": " + user.Email)
		return &FaultError{UserFaultErr, msg}
	}

	interactor.Logger.Log("Attempting to create user: " + user.Email)
	err := interactor.UserRepository.Create(user)
	return interactor.errorHandler(err)
}

func (interactor *UserInteractor) RemoveByEmail(userEmail string) MPError {
	err := interactor.UserRepository.Delete(userEmail)
	return interactor.errorHandler(err)
}

func (interactor *UserInteractor) GetByEmail(userEmail string) (*User, MPError) {
	interactor.Logger.Log("Getting user: " + userEmail)
	user, oneErr := interactor.UserRepository.One(userEmail)

	if oneErr != nil && oneErr.Error() == "not found" {
		interactor.Logger.Log("User not found: " + userEmail)
		return nil, &FaultError{UserFaultErr, oneErr.Error()}
	}

	return user, interactor.errorHandler(oneErr)
}

func (interactor *UserInteractor) GetAll() ([]User, MPError) {
	users, allErr := interactor.UserRepository.All()
	return users, interactor.errorHandler(allErr)
}
