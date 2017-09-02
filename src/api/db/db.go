package db

import (
	"gopkg.in/mgo.v2"
	"os"
)

const (
	dbNameKey        = "MP_DBNAME_ENV"
	userTableNameKey = "users"
)

// table holds the DB-table's configuration.
type table struct {
	Name string
}

// DB contains the current database session
type DB struct {
	Name      string
	Session   *mgo.Session
	UserTable table
}

// InitDB helps initialize the DB struct.
func InitDB(session *mgo.Session) *DB {
	return &DB{
		Name:    os.Getenv(dbNameKey),
		Session: session,
		UserTable: table{
			Name: userTableNameKey,
		},
	}
}
