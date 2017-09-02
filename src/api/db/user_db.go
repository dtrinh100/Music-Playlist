package db

import (
	"log"

	"github.com/dtrinh100/Music-Playlist/src/api/model"
)

// InsertUser inserts a user into the database
func (db *DB) InsertUser(username string, hashedPassword []byte, email string) error {
	conn := db.Session.DB(db.Name).C(db.UserTable.Name)
	insertErr := conn.Insert(&model.User{
		username,
		"",
		hashedPassword,
		email})
	if insertErr != nil {
		log.Print(insertErr)
		return insertErr
	}

	return nil
}
