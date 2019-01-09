package shareDB

import "vue-admin/web_server/core/mongo"

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
	ToJson() string

	isExist(query interface{}) bool
	insert(docs ...interface{}) error
	update(selector, update interface{}) error
	findOne(query, selector interface{}) (interface{}, error)
	findAll(query, selector interface{}) (results []interface{}, err error)
	remove(selector interface{}) error
	removeAll(selector interface{}) error
	totalCount(query, selector interface{}) (int, error)
	findPage(page, limit int, query, selector interface{}, fields ...string) (results []interface{}, err error)

	TotalCount(query, selector interface{}) (int, error)
	FindPage(page, limit int, query, selector interface{}, fields ...string) (results []interface{}, err error)
	Insert() error
	Update() error
	FindOne(query, selector interface{}) (interface{}, error)
	Remove() error
	RemoveAll() error
}
