package db

import (
	"log"

	"github.com/dtrinh100/Music-Playlist/src/api/model"
	"gopkg.in/mgo.v2"
)

// DB contains the current database session
type DB struct {
	session *mgo.Session
}

// InsertUser inserts a user into the database
func (d *DB) InsertUser(username string, password []byte, email string) error {
	conn := d.session.DB("mp").C("users")
	err := conn.Insert(&model.User{username, password, email})
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
