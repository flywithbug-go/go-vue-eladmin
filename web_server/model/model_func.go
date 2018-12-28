package model

import "doc-manager/web_server/core/mongo"

var db = "doc_manager"

func SetDBName(dbName string) {
	db = dbName
	mongo.SetIncrementDBName(dbName)
}

func init() {
	mongo.SetIncrementDBName(db)
}

func DBName() string {
	return db
}

type OperationModel interface {
	insert(docs ...interface{}) error
	isExist(query interface{}) bool
	findOne(query, selector interface{}) (*interface{}, error)
	findAll(query, selector interface{}) (results *[]interface{}, err error)
	totalCount(query, selector interface{}) (int, error)
	update(selector, update interface{}) error
	remove(selector interface{}) error
	removeAll(selector interface{}) error
}
