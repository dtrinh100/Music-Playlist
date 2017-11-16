package usecases

import "errors"

type UserInteractor struct {
  UserRepository UserRepository
  Logger         Logger
}

func (interactor *UserInteractor) UpdatePasswordForEmail(userEmail, password string) error {
  return interactor.UserRepository.Update(userEmail, M{"password": password})
}

func (interactor *UserInteractor) UpdateFirstNameForEmail(userEmail, firstName string) error {
  return interactor.UserRepository.Update(userEmail, M{"firstname": firstName})
}

func (interactor *UserInteractor) UpdateLastNameForEmail(userEmail, lastName string) error {
  return interactor.UserRepository.Update(userEmail, M{"lastname": lastName})
}

func (interactor *UserInteractor) UpdatePicURLForEmail(userEmail, picURL string) error {
  return interactor.UserRepository.Update(userEmail, M{"picurl": picURL})
}

// TODO: change contributions type to an array of songs?
// TODO: implement
func (interactor *UserInteractor) UpdateContributionsForEmail(userEmail string, contributions []interface{}) error {
  return nil
}

// TODO: change playlist type to an array of songs?
// TODO: implement
func (interactor *UserInteractor) UpdatePlaylistForEmail(userEmail string, playlist []interface{}) error {
  return nil
}

func (interactor *UserInteractor) CreateNew(user *User) error {
  if existingUser, oneErr := interactor.GetByEmail(user.Email); oneErr == nil && existingUser != nil {
    existsErr := errors.New("User already exists")
    interactor.Logger.Log(existsErr.Error() + ": " + user.Email)
    return existsErr
  }

  interactor.Logger.Log("Creating user: " + user.Email)
  return interactor.UserRepository.Create(user)
}

func (interactor *UserInteractor) RemoveByEmail(userEmail string) error {
  return interactor.UserRepository.Delete(userEmail)
}

func (interactor *UserInteractor) GetByEmail(userEmail string) (*User, error) {
  user, oneErr := interactor.UserRepository.One(userEmail)
  return user, oneErr
}

func (interactor *UserInteractor) GetAll() ([]User, error) {
  return interactor.UserRepository.All()
}
