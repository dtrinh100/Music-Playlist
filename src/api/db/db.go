package db

import "gopkg.in/mgo.v2"

// DB contains the current database session
type DB struct {
	Session *mgo.Session
}
