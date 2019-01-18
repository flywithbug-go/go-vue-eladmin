package model_monitor

import (
	"encoding/json"
	"time"
	"vue-admin/web_server/core/mongo"

	"gopkg.in/mgo.v2/bson"
)

const (
	logCollection = "log"
	dbName        = "monitor"
)

type Log struct {
	Time         string        `json:"time,omitempty" bson:"time,omitempty"`
	Code         string        `json:"code,omitempty" bson:"code,omitempty"`
	Info         string        `json:"info,omitempty" bson:"info,omitempty"`
	Level        int           `json:"level,omitempty" bson:"level,omitempty"`
	Flag         string        `json:"flag,omitempty" bson:"flag,omitempty"`
	ClientIp     string        `json:"client_ip,omitempty" bson:"client_ip,omitempty"`
	Method       string        `json:"method,omitempty" bson:"method,omitempty"`
	Path         string        `json:"path,omitempty" bson:"path,omitempty"`
	RequestId    string        `json:"request_id,omitempty" bson:"request_id,omitempty"`
	Latency      time.Duration `json:"latency,omitempty" bson:"latency,omitempty"`
	StatusCode   int           `json:"status_code,omitempty" bson:"status_code,omitempty"`
	UserId       int64         `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Para         interface{}   `json:"para,omitempty" bson:"para,omitempty"`
	ResponseCode int           `json:"response_code,omitempty" bson:"response_code,omitempty"`
	StartTime    int64         `json:"start_time,omitempty" bson:"start_time,omitempty"`
	EndTime      int64         `json:"end_time,omitempty" bson:"end_time,omitempty"`
	UUID         string        `json:"uuid,omitempty" bson:"uuid,omitempty"`
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

func (l Log) Update() error {
	if !l.isExist(bson.M{"request_id": l.RequestId}) {
		return l.Insert()
	}
	return l.update(bson.M{"request_id": l.RequestId}, l)
}

func (l Log) Insert() error {
	if l.isExist(bson.M{"request_id": l.RequestId}) {
		return l.Update()
	}
	if l.Para != nil {
		js, _ := json.Marshal(l.Para)
		if js != nil {
			l.Para = string(js)
		}
	}
	return l.insert(l)
}

func (l Log) TotalCount(query, selector interface{}) (int, error) {
	return l.totalCount(query, selector)
}
func (l Log) FindPageFilter(page, limit int, query, selector interface{}, fields ...string) ([]Log, error) {
	return l.findPage(page, limit, query, selector, fields...)
}
