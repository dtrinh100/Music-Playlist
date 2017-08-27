package db

import (
	"log"

	"github.com/dtrinh100/Music-Playlist/src/api/model"
	"github.com/dtrinh100/Music-Playlist/src/api/common"
)

// InsertUser inserts a user into the database
func (d *DB) InsertUser(username string, hashedpassword []byte, email string) error {
	conn := d.Session.DB(common.AppConfig.DB.Name).C(common.AppConfig.DB.UserTable.Name)
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
