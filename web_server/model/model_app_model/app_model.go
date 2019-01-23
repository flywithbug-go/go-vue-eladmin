package model_app_model

import (
	"encoding/json"
	"time"
	"vue-admin/web_server/core/mongo"
	"vue-admin/web_server/model/a_mongo_index"
	"vue-admin/web_server/model/shareDB"

	"gopkg.in/mgo.v2/bson"
)

const (
	AppModelCollection = mongo_index.CollectionAppModel
)

type AppModel struct {
	Id         int64 `json:"id,omitempty" bson:"_id,omitempty"`
	ModelId    int64 `json:"model_id,omitempty" bson:"model_id,omitempty"`
	AppId      int64 `json:"app_id,omitempty" bson:"app_id,omitempty"`
	CreateTime int64 `json:"create_time,omitempty" bson:"create_time,omitempty"`
}

func (r AppModel) ToJson() string {
	js, _ := json.Marshal(r)
	return string(js)
}

func (r AppModel) isExist(query interface{}) bool {
	return mongo.IsExist(shareDB.DocManagerDBName(), AppModelCollection, query)
}

func (r AppModel) insert(docs ...interface{}) error {
	return mongo.Insert(shareDB.DocManagerDBName(), AppModelCollection, docs...)
}

func (r AppModel) update(selector, update interface{}) error {
	return mongo.Update(shareDB.DocManagerDBName(), AppModelCollection, selector, update, true)
}

func (r AppModel) findOne(query, selector interface{}) (AppModel, error) {
	ap := AppModel{}
	err := mongo.FindOne(shareDB.DocManagerDBName(), AppModelCollection, query, selector, &ap)
	return ap, err
}
func (r AppModel) findAll(query, selector interface{}) (results []AppModel, err error) {
	results = []AppModel{}
	err = mongo.FindAll(shareDB.DocManagerDBName(), AppModelCollection, query, selector, &results)
	return results, err
}

func (r AppModel) remove(selector interface{}) error {
	return mongo.Remove(shareDB.DocManagerDBName(), AppModelCollection, selector)
}

func (r AppModel) removeAll(selector interface{}) error {
	return mongo.RemoveAll(shareDB.DocManagerDBName(), AppModelCollection, selector)
}

func (r AppModel) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(shareDB.DocManagerDBName(), AppModelCollection, query, selector)
}

func (r AppModel) findPage(page, limit int, query, selector interface{}, fields ...string) (results []AppModel, err error) {
	results = []AppModel{}
	err = mongo.FindPage(shareDB.DocManagerDBName(), AppModelCollection, page, limit, query, selector, &results, fields...)
	return
}

func (r AppModel) FindOne(query, selector interface{}) (role AppModel, err error) {
	role, err = r.findOne(query, selector)
	return
}
func (r AppModel) FindAll(query, selector interface{}) (results []AppModel, err error) {
	results = []AppModel{}
	return r.findAll(query, selector)
}

func (r AppModel) Exist(query interface{}) bool {
	return r.isExist(query)
}

func (r AppModel) Insert() error {
	r.Id, _ = mongo.GetIncrementId(shareDB.DocManagerDBName(), AppModelCollection)
	r.CreateTime = time.Now().Unix()
	return r.insert(r)
}

func (r AppModel) Update() error {
	return r.update(bson.M{"_id": r.Id}, r)
}

func (r AppModel) Remove() error {
	return r.remove(bson.M{"_id": r.Id})
}

func (r AppModel) RemoveModelId(modelId int64) error {
	return r.removeAll(bson.M{"model_id": modelId})
}

func (r AppModel) RemoveAppId(appId int64) error {
	return r.removeAll(bson.M{"app_id": appId})
}
