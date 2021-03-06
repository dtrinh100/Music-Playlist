package infrastructure

import (
	"gopkg.in/mgo.v2"
	"errors"
	"github.com/dtrinh100/Music-Playlist/src/api/usecases"
)

type MongoHandler struct {
	session     *mgo.Session
	dbName      string
	dbTableName string
}

func (handler *MongoHandler) col() *mgo.Collection {
	return handler.session.DB(handler.dbName).C(handler.dbTableName)
}

func (handler *MongoHandler) EnsureIndex(fieldName string) error {
	dbIndex := mgo.Index{
		Key: []string{fieldName},
		Unique: true,
		Background: true,
		Sparse: false,
	}

	return handler.col().EnsureIndex(dbIndex)
}

func (handler *MongoHandler) FindAndModify(query, update usecases.M, result interface{}) (interface{}, error) {
	change := mgo.Change{
		Update:    update,
		ReturnNew: true,
	}

	_, changeErr := handler.col().Find(query).Apply(change, result)

	return change, changeErr
}

func (handler *MongoHandler) Create(docs ...interface{}) error {
	return handler.col().Insert(docs...)
}

func (handler *MongoHandler) Update(selector usecases.M, update usecases.M) error {
	return handler.col().Update(selector, usecases.M{"$set": update})
}

func (handler *MongoHandler) Delete(selector usecases.M) error {
	return handler.col().Remove(selector)
}

func (handler *MongoHandler) One(query usecases.M, result interface{}) error {
	qry := handler.col().Find(query)
	if qry == nil {
		return errors.New("failed to obtain a query for One()")
	}

	return qry.One(result)
}

func (handler *MongoHandler) All(results interface{}) error {
	var qry *mgo.Query = handler.col().Find(nil)
	if qry == nil {
		return errors.New("failed to obtain a query for All()")
	}

	return qry.All(results)
}
