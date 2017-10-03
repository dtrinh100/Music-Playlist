package db

import (
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"time"
)

const (
	dbNameKey        = "MP_DBNAME_ENV"
	userTableNameKey = "users"
)

// Note: 'MPDatabase' name comes from docker-compose.yml
const dbURLAddress = "MPDatabase"

// table holds the DB-table's configuration.
type table struct {
	Name string
}

// DB contains the current database session
type DB struct {
	Name      string
	session   *mgo.Session
	UserTable table
}

// GetSessionCopy returns a copy of the main session.
func (db *DB) GetSessionCopy() *mgo.Session {
	if db.session == nil {
		db.createDBSession()
	}

	return db.session.Copy()
}

// createDBSession initializes MongoDB session
func (db *DB) createDBSession() {
	// NOTE: this main session closes only when the entire app shuts down.
	//		 It will be used as the original, to make copies.
	session, sessionErr := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{dbURLAddress},
		Username: "",
		Password: "",
		Timeout:  60 * time.Second,
	})

	if sessionErr != nil {
		log.Fatal("Failed to obtain a DB session"+":", sessionErr)
	}

	db.session = session
}

// addIndexes adds indexes into MongoDB
func (db *DB) addIndexes() {
	userIndex := mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}

	session := db.GetSessionCopy()
	defer session.Close()

	userCol := session.DB(db.Name).C(db.UserTable.Name)

	ensureErr := userCol.EnsureIndex(userIndex)
	if ensureErr != nil {
		log.Fatal("Failed to EnsureIndex for "+db.UserTable.Name+":", ensureErr)
	}

}

// InitDB helps initialize the DB struct.
func InitDB() *DB {
	db := &DB{
		Name: os.Getenv(dbNameKey),
		UserTable: table{
			Name: userTableNameKey,
		},
		session: nil,
	}

	db.createDBSession()
	db.addIndexes()

	return db
}
