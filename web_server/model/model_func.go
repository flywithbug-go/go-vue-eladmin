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
	Insert(docs ...interface{}) error
	IsExist(query interface{}) bool
	FindOne(query, selector interface{}) (*interface{}, error)
	FindAll(query, selector interface{}, result *[]interface{}) error
	Update(selector, update interface{}) error
	Remove(selector interface{}) error
	RemoveAll(selector interface{}) error
}
