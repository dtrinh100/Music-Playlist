package db

import (
	"gopkg.in/mgo.v2"
	"os"
)

const (
	dbname_envkey  = "MP_DBNAME_ENV"
	usertable_name = "users"
)

type table struct {
	Name string
}

// DB contains the current database session
type DB struct {
	Name      string
	Session   *mgo.Session
	UserTable table
}

func InitDB(session *mgo.Session) *DB {
	return &DB{
		Name:    os.Getenv(dbname_envkey),
		Session: session,
		UserTable: table{
			Name: usertable_name,
		},
	}
}
