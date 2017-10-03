package db

import (
	"github.com/dtrinh100/Music-Playlist/src/api/model"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct {
	DBName, Name string
	Session      *mgo.Session
}

func (userRepo *UserRepository) Verify(user *model.User) error {
	fields := bson.M{
		"_id":      0,
		"email":    1,
		"username": 1}

	if queryErr := userRepo.collection().Find(bson.M{
		"email": user.Email}).Select(fields).One(&user); queryErr != nil {
		return queryErr
	}

	return nil
}

func (userRepo *UserRepository) Login(user *model.User) error {
	fields := bson.M{
		"_id":            0,
		"username":       1,
		"email":          1,
		"hashedpassword": 1}

	pw := user.Password

	if queryErr := userRepo.collection().Find(bson.M{
		"email": user.Email}).Select(fields).One(&user); queryErr != nil {
		return queryErr
	}

	if pwErr := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(pw)); pwErr != nil {
		return pwErr
	}

	return nil
}

func (userRepo *UserRepository) Register(user *model.User) error {
	// Encrypts the password of the user before storing it in the db
	// recommended use a cost of 12 or more
	hashedPassword, genErr := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if genErr != nil {
		return genErr
	}

	user.HashedPassword = hashedPassword

	if insertErr := userRepo.collection().Insert(user); insertErr != nil {
		return insertErr
	}

	return nil
}

func (userRepo *UserRepository) collection() *mgo.Collection {
	return userRepo.Session.DB(userRepo.DBName).C(userRepo.Name)
}
