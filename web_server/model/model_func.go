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
	FindAll() ([]interface{}, error)
	Insert() error
	Update(id string) error
	Remove(id string) error
	FindById(id string) (interface{}, error)
}
