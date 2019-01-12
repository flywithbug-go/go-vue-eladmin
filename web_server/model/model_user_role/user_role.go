package model_user_role

import (
	"encoding/json"
	"time"
	"vue-admin/web_server/core/mongo"
	"vue-admin/web_server/model/mongo_index"
	"vue-admin/web_server/model/shareDB"

	"gopkg.in/mgo.v2/bson"
)

const (
	UserRoleCollection = mongo_index.CollectionUserRole
)

type UserRole struct {
	Id         int64 `json:"id,omitempty" bson:"_id,omitempty"`
	UserId     int64 `json:"user_id" bson:"user_id"`
	RoleId     int64 `json:"role_id" bson:"role_id"`
	CreateTime int64 `json:"create_time" bson:"create_time"`
}

func (r UserRole) ToJson() string {
	js, _ := json.Marshal(r)
	return string(js)
}

func (r UserRole) isExist(query interface{}) bool {
	return mongo.IsExist(shareDB.DBName(), UserRoleCollection, query)
}

func (r UserRole) insert(docs ...interface{}) error {
	return mongo.Insert(shareDB.DBName(), UserRoleCollection, docs...)
}

func (r UserRole) update(selector, update interface{}) error {
	return mongo.Update(shareDB.DBName(), UserRoleCollection, selector, update, true)
}

func (r UserRole) findOne(query, selector interface{}) (UserRole, error) {
	ap := UserRole{}
	err := mongo.FindOne(shareDB.DBName(), UserRoleCollection, query, selector, &ap)
	return ap, err
}
func (r UserRole) findAll(query, selector interface{}) (results []UserRole, err error) {
	results = []UserRole{}
	err = mongo.FindAll(shareDB.DBName(), UserRoleCollection, query, selector, &results)
	return results, err
}

func (r UserRole) remove(selector interface{}) error {
	return mongo.Remove(shareDB.DBName(), UserRoleCollection, selector)
}

func (r UserRole) removeAll(selector interface{}) error {
	return mongo.RemoveAll(shareDB.DBName(), UserRoleCollection, selector)
}

func (r UserRole) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(shareDB.DBName(), UserRoleCollection, query, selector)
}

func (r UserRole) findPage(page, limit int, query, selector interface{}, fields ...string) (results []UserRole, err error) {
	results = []UserRole{}
	err = mongo.FindPage(shareDB.DBName(), UserRoleCollection, page, limit, query, selector, &results, fields...)
	return
}

func (r UserRole) FindOne(query, selector interface{}) (role UserRole, err error) {
	role, err = r.findOne(query, selector)
	return
}
func (r UserRole) FindAll(query, selector interface{}) (results []UserRole, err error) {
	results = []UserRole{}
	return r.findAll(query, selector)
}

func (r UserRole) Exist(query interface{}) bool {
	return r.isExist(query)
}

func (r UserRole) Insert() error {
	r.Id, _ = mongo.GetIncrementId(UserRoleCollection)
	r.CreateTime = time.Now().Unix() * 1000
	return r.insert(r)
}

func (r UserRole) Update() error {
	return r.update(bson.M{"_id": r.Id}, r)
}

func (r UserRole) Remove() error {
	return r.remove(bson.M{"_id": r.Id})
}

func (r UserRole) RemoveUserId(userId int64) error {
	return r.removeAll(bson.M{"user_id": userId})
}
