package interfaces

import (
	"github.com/dtrinh100/Music-Playlist/src/api/usecases"
	"golang.org/x/crypto/bcrypt"
)

type DBRepo struct {
	dbHandlers map[string]DBHandler
	dbHandler  DBHandler
}

type DBUserRepo DBRepo

func (repo *DBUserRepo) Create(user *usecases.User) error {
	hashedPass, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		return hashErr
	}

	user.HashedPassword = hashedPass
	return repo.dbHandler.Create(*user)
}

func (repo *DBUserRepo) One(userEmail string) (*usecases.User, error) {
	user := new(usecases.User)
	oneErr := repo.dbHandler.One(usecases.M{"email": userEmail}, user)

	return user, oneErr
}

func (repo *DBUserRepo) All() ([]usecases.User, error) {
	users := []usecases.User{}
	allErr := repo.dbHandler.All(&users)

	return users, allErr
}

func (repo *DBUserRepo) Update(userEmail string, changes usecases.M) error {
	return repo.dbHandler.Update(usecases.M{"email": userEmail}, changes)
}

func (repo *DBUserRepo) Delete(userEmail string) error {
	return repo.dbHandler.Delete(usecases.M{"email": userEmail})
}

func NewDBUserRepo(dbHandlers map[string]DBHandler) *DBUserRepo {
	dbUserRepo := new(DBUserRepo)
	dbUserRepo.dbHandlers = dbHandlers
	dbUserRepo.dbHandler = dbHandlers["DBUserRepo"]

	return dbUserRepo
}
