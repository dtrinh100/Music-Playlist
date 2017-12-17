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
	interactor.Logger.Log("Getting user by email: " + userEmail)
	user, oneErr := interactor.UserRepository.OneByEmail(userEmail)

	if oneErr != nil && oneErr.Error() == "not found" {
		interactor.Logger.Log("User email not found: " + userEmail)
		return nil, &FaultError{UserFaultErr, oneErr.Error()}
	}

	return user, interactor.errorHandler(oneErr)
}

func (interactor *UserInteractor) GetByUsername(username string) (*User, MPError) {
	interactor.Logger.Log("Getting user by username: " + username)
	user, oneErr := interactor.UserRepository.OneByUsername(username)

	if oneErr != nil && oneErr.Error() == "not found" {
		interactor.Logger.Log("Username not found: " + username)
		return nil, &FaultError{UserFaultErr, oneErr.Error()}
	}

	return user, interactor.errorHandler(oneErr)
}

func (interactor *UserInteractor) GetAll() ([]User, MPError) {
	users, allErr := interactor.UserRepository.All()
	return users, interactor.errorHandler(allErr)
}

func (interactor *UserInteractor) ComparePassword(user *User, clearTextPass string) MPError {
	compareErr := interactor.UserRepository.ComparePassword(user.HashedPassword, []byte(clearTextPass))

	if compareErr != nil && compareErr.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password" {
		interactor.Logger.Log("Invalid credentials: " + user.Email)
		return &FaultError{UserFaultErr, "invalid credentials"}
	}

	return interactor.errorHandler(compareErr)
}
