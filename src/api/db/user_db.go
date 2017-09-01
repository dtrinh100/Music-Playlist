package db

import (
	"log"

	"github.com/dtrinh100/Music-Playlist/src/api/model"
)

// InsertUser inserts a user into the database
func (db *DB) InsertUser(username string, hashedpassword []byte, email string) error {
	conn := db.Session.DB(db.Name).C(db.UserTable.Name)
	err := conn.Insert(&model.User{
		username,
		"",
		hashedpassword,
		email})
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
