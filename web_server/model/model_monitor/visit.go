package model_monitor

import (
	"encoding/json"
	"fmt"
	"vue-admin/web_server/core/mongo"
	"vue-admin/web_server/model/a_mongo_index"
	"vue-admin/web_server/model/shareDB"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	visitCollection = mongo_index.CollectionVisit
)

type Visit struct {
	ClientIp string `json:"client_ip,omitempty" bson:"client_ip,omitempty"`
	UUID     string `json:"uuid,omitempty" bson:"uuid,omitempty"`
	Count    int64  `json:"count,omitempty" bson:"count,omitempty"`         //访问次数
	TimeDate string `json:"time_date,omitempty" bson:"time_date,omitempty"` //2018-06-10
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
	if len(v.ClientIp) == 0 || len(v.UUID) == 0 {
		return fmt.Errorf("client_ip or uuid is null")
	}
	return v.insert(v)
}

func (v Visit) IncrementVisit() (int64, error) {
	if len(v.ClientIp) == 0 || len(v.UUID) == 0 {
		return -1, fmt.Errorf("client_ip or uuid is null")
	}
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"count": 1}, "$set": bson.M{"time_date": v.TimeDate}},
		ReturnNew: true,
	}
	_, c := mongo.Collection(shareDB.MonitorDBName(), visitCollection)
	_, err := c.Find(bson.M{"client_ip": v.ClientIp, "uuid": v.UUID}).Apply(change, v)
	if err != nil {
		v.Count = 1
		err = v.Insert()
		if err != nil {
			return -1, err
		}
	}
	return v.Count, nil
}

func (v Visit) TotalCount(query, selector interface{}) (int, error) {
	return v.totalCount(query, selector)
}
func (v Visit) FindPageFilter(page, limit int, query, selector interface{}, fields ...string) ([]Visit, error) {
	return v.findPage(page, limit, query, selector, fields...)
}
