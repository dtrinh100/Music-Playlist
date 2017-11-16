package infrastructure

import (
  "gopkg.in/mgo.v2"
  "time"
  "errors"
  "github.com/dtrinh100/Music-Playlist/src/api/usecases"
)

type MongoHandler struct {
  session     *mgo.Session
  dbName      string
  dbTableName string
}

func (handler *MongoHandler) Create(docs ...interface{}) error {
  return handler.session.DB(handler.dbName).C(handler.dbTableName).Insert(docs...)
}

func (handler *MongoHandler) Update(selector usecases.M, update usecases.M) error {
  return handler.session.DB(handler.dbName).C(handler.dbTableName).Update(selector, usecases.M{"$set": update})
}

func (handler *MongoHandler) Delete(selector usecases.M) error {
  return handler.session.DB(handler.dbName).C(handler.dbTableName).Remove(selector)
}

func (handler *MongoHandler) One(query usecases.M, result interface{}) error {
  qry := handler.session.DB(handler.dbName).C(handler.dbTableName).Find(query)
  if qry == nil {
    return errors.New("failed to obtain a query for One()")
  }

  return qry.One(result)
}

func (handler *MongoHandler) All(results interface{}) error {
  var qry *mgo.Query = handler.session.DB(handler.dbName).C(handler.dbTableName).Find(nil)
  if qry == nil {
    return errors.New("failed to obtain a query for All()")
  }

  return qry.All(results)
}

func NewMongoHandler(session *mgo.Session, dbName, dbTableName string) *MongoHandler {
  mongoHandler := new(MongoHandler)
  mongoHandler.session = session
  mongoHandler.dbName = dbName
  mongoHandler.dbTableName = dbTableName

  return mongoHandler
}

func NewMongoSession(addrs, un, pw string) *mgo.Session {
  session, sessionErr := mgo.DialWithInfo(&mgo.DialInfo{
    Addrs:    []string{addrs},
    Username: un,
    Password: pw,
    Timeout:  60 * time.Second,
  })

  if sessionErr != nil {
    return nil // TODO: throw a fatal-error?
  }
  return session
}
