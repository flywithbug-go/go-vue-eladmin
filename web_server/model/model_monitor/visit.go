package model_monitor

import (
	"encoding/json"
	"fmt"
	"vue-admin/web_server/core/mongo"
)

type Visit struct {
	ClientIp string `json:"client_ip,omitempty" bson:"client_ip,omitempty"`
	UUID     int64  `json:"uuid,omitempty" bson:"uuid,omitempty"`
}

func (v Visit) ToJson() string {
	js, _ := json.Marshal(v)
	return string(js)
}

func (v Visit) isExist(query interface{}) bool {
	return mongo.IsExist(dbName, logCollection, query)
}

func (v Visit) insert(docs ...interface{}) error {
	return mongo.Insert(dbName, logCollection, docs...)
}

func (v Visit) update(selector, update interface{}) error {
	return mongo.Update(dbName, logCollection, selector, update, true)
}

func (v Visit) findOne(query, selector interface{}) (Log, error) {
	ap := Log{}
	err := mongo.FindOne(dbName, logCollection, query, selector, &ap)
	return ap, err
}
func (v Visit) findAll(query, selector interface{}) (results []Log, err error) {
	results = []Log{}
	err = mongo.FindAll(dbName, logCollection, query, selector, &results)
	return results, err
}

func (v Visit) remove(selector interface{}) error {
	return mongo.Remove(dbName, logCollection, selector)
}

func (v Visit) removeAll(selector interface{}) error {
	return mongo.RemoveAll(dbName, logCollection, selector)
}

func (v Visit) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(dbName, logCollection, query, selector)
}

func (v Visit) findPage(page, limit int, query, selector interface{}, fields ...string) (results []Visit, err error) {
	results = []Visit{}
	err = mongo.FindPage(dbName, logCollection, page, limit, query, selector, &results, fields...)
	return
}

func (v Visit) pipeAll(pipeline, result interface{}, allowDiskUse bool) error {
	return mongo.PipeAll(dbName, logCollection, pipeline, result, allowDiskUse)
}

func (v Visit) pipeOne(pipeline, result interface{}, allowDiskUse bool) error {
	return mongo.PipeOne(dbName, logCollection, pipeline, result, allowDiskUse)
}

func (v Visit) explain(pipeline, result interface{}) (results []Visit, err error) {
	err = mongo.Explain(dbName, logCollection, pipeline, result)
	return
}

func (v Visit) Insert() error {
	if len(v.ClientIp) == 0 {
		return fmt.Errorf("client_ip is null")
	}
	return v.insert(v)
}

func (v Visit) TotalCount(query, selector interface{}) (int, error) {
	return v.totalCount(query, selector)
}
func (v Visit) FindPageFilter(page, limit int, query, selector interface{}, fields ...string) ([]Visit, error) {
	return v.findPage(page, limit, query, selector, fields...)
}
