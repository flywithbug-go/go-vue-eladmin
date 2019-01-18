package model_monitor

import (
	"encoding/json"
	"fmt"
	"vue-admin/web_server/core/mongo"
	"vue-admin/web_server/model/a_mongo_index"
	"vue-admin/web_server/model/shareDB"
)

const (
	visitCollection = mongo_index.CollectionVisit
)

type Visit struct {
	ClientIp string `json:"client_ip,omitempty" bson:"client_ip,omitempty"`
	UUID     int64  `json:"uuid,omitempty" bson:"uuid,omitempty"`
	Count    int    `json:"count,omitempty" bson:"count,omitempty"` //访问次数
}

func (v Visit) ToJson() string {
	js, _ := json.Marshal(v)
	return string(js)
}

func (v Visit) isExist(query interface{}) bool {
	return mongo.IsExist(shareDB.MonitorDBName(), visitCollection, query)
}

func (v Visit) insert(docs ...interface{}) error {
	return mongo.Insert(shareDB.MonitorDBName(), visitCollection, docs...)
}

func (v Visit) update(selector, update interface{}) error {
	return mongo.Update(shareDB.MonitorDBName(), visitCollection, selector, update, true)
}

func (v Visit) findOne(query, selector interface{}) (Log, error) {
	ap := Log{}
	err := mongo.FindOne(shareDB.MonitorDBName(), visitCollection, query, selector, &ap)
	return ap, err
}
func (v Visit) findAll(query, selector interface{}) (results []Log, err error) {
	results = []Log{}
	err = mongo.FindAll(shareDB.MonitorDBName(), visitCollection, query, selector, &results)
	return results, err
}

func (v Visit) remove(selector interface{}) error {
	return mongo.Remove(shareDB.MonitorDBName(), visitCollection, selector)
}

func (v Visit) removeAll(selector interface{}) error {
	return mongo.RemoveAll(shareDB.MonitorDBName(), visitCollection, selector)
}

func (v Visit) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(shareDB.MonitorDBName(), visitCollection, query, selector)
}

func (v Visit) findPage(page, limit int, query, selector interface{}, fields ...string) (results []Visit, err error) {
	results = []Visit{}
	err = mongo.FindPage(shareDB.MonitorDBName(), visitCollection, page, limit, query, selector, &results, fields...)
	return
}

func (v Visit) pipeAll(pipeline, result interface{}, allowDiskUse bool) error {
	return mongo.PipeAll(shareDB.MonitorDBName(), visitCollection, pipeline, result, allowDiskUse)
}

func (v Visit) pipeOne(pipeline, result interface{}, allowDiskUse bool) error {
	return mongo.PipeOne(shareDB.MonitorDBName(), visitCollection, pipeline, result, allowDiskUse)
}

func (v Visit) explain(pipeline, result interface{}) (results []Visit, err error) {
	err = mongo.Explain(shareDB.MonitorDBName(), visitCollection, pipeline, result)
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
