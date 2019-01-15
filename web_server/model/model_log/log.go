package model_log

import (
	"encoding/json"
	"vue-admin/web_server/core/mongo"
	"vue-admin/web_server/model/shareDB"

	"gopkg.in/mgo.v2/bson"
)

const (
	logCollection = "log_info"
	dbName        = "log"
)

type Log struct {
	Id    int64  `json:"id,omitempty" bson:"_id,omitempty"`
	Time  string `json:"time,omitempty" bson:"time,omitempty"`
	Code  string `json:"code,omitempty" bson:"code,omitempty"`
	Info  string `json:"info,omitempty" bson:"info,omitempty"`
	Level int    `json:"level,omitempty" bson:"level,omitempty"`
}

func (l Log) ToJson() string {
	js, _ := json.Marshal(l)
	return string(js)
}

func (l Log) isExist(query interface{}) bool {
	return mongo.IsExist(dbName, logCollection, query)
}

func (l Log) insert(docs ...interface{}) error {
	return mongo.Insert(dbName, logCollection, docs...)
}

func (l Log) update(selector, update interface{}) error {
	return mongo.Update(dbName, logCollection, selector, update, true)
}

func (l Log) findOne(query, selector interface{}) (Log, error) {
	ap := Log{}
	err := mongo.FindOne(dbName, logCollection, query, selector, &ap)
	return ap, err
}
func (l Log) findAll(query, selector interface{}) (results []Log, err error) {
	results = []Log{}
	err = mongo.FindAll(dbName, logCollection, query, selector, &results)
	return results, err
}

func (l Log) remove(selector interface{}) error {
	return mongo.Remove(dbName, logCollection, selector)
}

func (l Log) removeAll(selector interface{}) error {
	return mongo.RemoveAll(dbName, logCollection, selector)
}

func (l Log) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(dbName, logCollection, query, selector)
}

func (l Log) findPage(page, limit int, query, selector interface{}, fields ...string) (results []Log, err error) {
	results = []Log{}
	err = mongo.FindPage(dbName, logCollection, page, limit, query, selector, &results, fields...)
	return
}

func (l Log) pipeAll(pipeline, result interface{}, allowDiskUse bool) error {
	return mongo.PipeAll(dbName, logCollection, pipeline, result, allowDiskUse)
}

func (l Log) pipeOne(pipeline, result interface{}, allowDiskUse bool) error {
	return mongo.PipeOne(dbName, logCollection, pipeline, result, allowDiskUse)
}

func (l Log) explain(pipeline, result interface{}) (results []Log, err error) {
	err = mongo.Explain(dbName, logCollection, pipeline, result)
	return
}

func (l Log) Exist() bool {
	return l.isExist(bson.M{"_id": l.Id})
}

func (l Log) Insert() (int64, error) {
	l.Id, _ = mongo.GetIncrementId(shareDB.DBName(), logCollection)

	return -1, nil
}
