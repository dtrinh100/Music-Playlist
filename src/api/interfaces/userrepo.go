package interfaces

import (
	"github.com/dtrinh100/Music-Playlist/src/api/usecases"
	"golang.org/x/crypto/bcrypt"
	"errors"
)

type DBUserRepo struct {
	DBRepo
}

func (repo *DBUserRepo) Create(user *usecases.User) error {
	hashedPass, hashErr := repo.getHashedPass(user.Password)
	if hashErr != nil {
		return hashErr
	}

	user.HashedPassword = hashedPass

	if createErr := repo.dbHandler.Create(*user); createErr != nil {
		return createErr
	}

	return nil
}

func (repo *DBUserRepo) OneByEmail(userEmail string) (*usecases.User, error) {
	user := new(usecases.User)
	oneErr := repo.dbHandler.One(usecases.M{"email": userEmail}, user)

	return user, oneErr
}

func (repo *DBUserRepo) OneByUsername(username string) (*usecases.User, error) {
	user := new(usecases.User)
	oneErr := repo.dbHandler.One(usecases.M{"username": username}, user)

	return user, oneErr
}

func (repo *DBUserRepo) All() ([]usecases.User, error) {
	users := []usecases.User{}
	allErr := repo.dbHandler.All(&users)

	return users, allErr
}

func (repo *DBUserRepo) Update(userEmail string, changes usecases.M) error {
	if changes["password"] != nil {
		pass, strOk := changes["password"].(string)

		if !strOk {
			return errors.New("string conversion failed")
		}

		epass, hashErr := repo.getHashedPass(pass)

		if hashErr != nil {
			return errors.New("password hashing failed")
		}

		delete(changes, "password")
		changes["hashedpassword"] = epass
	}

	return repo.dbHandler.Update(usecases.M{"email": userEmail}, changes)
}

func (repo *DBUserRepo) Delete(userEmail string) error {
	return repo.dbHandler.Delete(usecases.M{"email": userEmail})
}

func (repo *DBUserRepo) ComparePassword(hashedPass, clearTextPass []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPass, clearTextPass)
}

func (repo *DBUserRepo) getHashedPass(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

func NewDBUserRepo(dbHandlers map[string]DBHandler) *DBUserRepo {
	dbUserRepo := new(DBUserRepo)
	dbUserRepo.dbHandlers = dbHandlers
	dbUserRepo.dbHandler = dbHandlers["DBUserRepo"]

	return dbUserRepo
}
