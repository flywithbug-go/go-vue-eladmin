package model_app_user

import (
	"encoding/json"
	"time"
	"vue-admin/web_server/core/mongo"
	"vue-admin/web_server/model/a_mongo_index"
	"vue-admin/web_server/model/shareDB"

	"gopkg.in/mgo.v2/bson"
)

const (
	AppUserCollection = mongo_index.CollectionAppUser
)

type AppUser struct {
	Id         int64 `json:"id,omitempty" bson:"_id,omitempty"`
	UserId     int64 `json:"user_id,omitempty" bson:"user_id,omitempty"`
	RoleId     int64 `json:"role_id,omitempty" bson:"role_id,omitempty"`
	CreateTime int64 `json:"create_time,omitempty" bson:"create_time,omitempty"`
}

func (r AppUser) ToJson() string {
	js, _ := json.Marshal(r)
	return string(js)
}

func (r AppUser) isExist(query interface{}) bool {
	return mongo.IsExist(shareDB.DocManagerDBName(), AppUserCollection, query)
}

func (r AppUser) insert(docs ...interface{}) error {
	return mongo.Insert(shareDB.DocManagerDBName(), AppUserCollection, docs...)
}

func (r AppUser) update(selector, update interface{}) error {
	return mongo.Update(shareDB.DocManagerDBName(), AppUserCollection, selector, update, true)
}

func (r AppUser) findOne(query, selector interface{}) (AppUser, error) {
	ap := AppUser{}
	err := mongo.FindOne(shareDB.DocManagerDBName(), AppUserCollection, query, selector, &ap)
	return ap, err
}
func (r AppUser) findAll(query, selector interface{}) (results []AppUser, err error) {
	results = []AppUser{}
	err = mongo.FindAll(shareDB.DocManagerDBName(), AppUserCollection, query, selector, &results)
	return results, err
}

func (r AppUser) remove(selector interface{}) error {
	return mongo.Remove(shareDB.DocManagerDBName(), AppUserCollection, selector)
}

func (r AppUser) removeAll(selector interface{}) error {
	return mongo.RemoveAll(shareDB.DocManagerDBName(), AppUserCollection, selector)
}

func (r AppUser) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(shareDB.DocManagerDBName(), AppUserCollection, query, selector)
}

func (r AppUser) findPage(page, limit int, query, selector interface{}, fields ...string) (results []AppUser, err error) {
	results = []AppUser{}
	err = mongo.FindPage(shareDB.DocManagerDBName(), AppUserCollection, page, limit, query, selector, &results, fields...)
	return
}

func (r AppUser) FindOne(query, selector interface{}) (role AppUser, err error) {
	role, err = r.findOne(query, selector)
	return
}
func (r AppUser) FindAll(query, selector interface{}) (results []AppUser, err error) {
	results = []AppUser{}
	return r.findAll(query, selector)
}

func (r AppUser) Exist(query interface{}) bool {
	return r.isExist(query)
}

func (r AppUser) Insert() error {
	r.Id, _ = mongo.GetIncrementId(shareDB.DocManagerDBName(), AppUserCollection)
	r.CreateTime = time.Now().Unix() * 1000
	return r.insert(r)
}

func (r AppUser) Update() error {
	return r.update(bson.M{"_id": r.Id}, r)
}

func (r AppUser) Remove() error {
	return r.remove(bson.M{"_id": r.Id})
}

func (r AppUser) RemoveUserId(userId int64) error {
	return r.removeAll(bson.M{"user_id": userId})
}
