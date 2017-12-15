package infrastructure

import (
	"github.com/dtrinh100/Music-Playlist/src/api/interfaces"
	"github.com/dtrinh100/Music-Playlist/src/api/usecases"
	"gopkg.in/mgo.v2"
	"time"
)

func newMongoHandler(session *mgo.Session, dbName, dbTableName string) *MongoHandler {
	mongoHandler := new(MongoHandler)
	mongoHandler.session = session
	mongoHandler.dbName = dbName
	mongoHandler.dbTableName = dbTableName

	return mongoHandler
}

func newMongoSession(addrs, un, pw string) *mgo.Session {
	session, sessionErr := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{addrs},
		Username: un,
		Password: pw,
		Timeout:  60 * time.Second,
	})

	if sessionErr != nil {
		panic(sessionErr)
		return nil
	}
	return session
}

func GetAndInitDBHandlersForDBName(dbName string) map[string]interfaces.DBHandler {
	session := newMongoSession("MPDatabase", "", "")
	dbUserHandler := newMongoHandler(session, dbName, "usertable")
	dbSongHandler := newMongoHandler(session, dbName, "songtable")
	dbCounterHandler := newMongoHandler(session, dbName, "countertable")

	songSeq := interfaces.Counter{"songid", 0}
	userSeq := interfaces.Counter{"userid", 0}
	dbCounterHandler.Create(songSeq)
	dbCounterHandler.Create(userSeq)

	if ensureErr := dbUserHandler.EnsureIndex("username"); ensureErr != nil {
		panic(ensureErr)
	}

	handlers := make(map[string]interfaces.DBHandler)
	handlers["DBUserRepo"] = dbUserHandler
	handlers["DBSongRepo"] = dbSongHandler
	handlers["DBCounterRepo"] = dbCounterHandler

	return handlers
}

func GetAndInitUserInteractor(logger *Logger, handlers map[string]interfaces.DBHandler) *usecases.UserInteractor {
	userInteractor := new(usecases.UserInteractor)
	userInteractor.UserRepository = interfaces.NewDBUserRepo(handlers)
	userInteractor.Logger = logger

	return userInteractor
}

func GetAndInitSongInteractor(logger *Logger, handlers map[string]interfaces.DBHandler) *usecases.SongInteractor {
	songInteractor := new(usecases.SongInteractor)
	songInteractor.SongRepository = interfaces.NewDBSongRepo(handlers)
	songInteractor.Logger = logger

	return songInteractor
}
