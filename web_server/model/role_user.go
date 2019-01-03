package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"vue-admin/web_server/core/mongo"

	"gopkg.in/mgo.v2/bson"
)

const (
	userRoleCollection = "user_role"
)

type UserRole struct {
	Id     int64 `json:"id,omitempty" bson:"_id,omitempty"`
	UserId int64 `json:"user_id,omitempty"  bson:"user_id,omitempty"`
	RoleId int64 `json:"role_id,omitempty" bson:"role_id,omitempty"`
}

func (r UserRole) ToJson() string {
	js, _ := json.Marshal(r)
	return string(js)
}

func (r UserRole) isExist(query interface{}) bool {
	return mongo.IsExist(db, roleCollection, query)
}

func (r UserRole) insert(docs ...interface{}) error {
	return mongo.Insert(db, roleCollection, docs...)
}

func (r UserRole) update(selector, update interface{}) error {
	return mongo.Update(db, roleCollection, selector, update, true)
}

func (r UserRole) findOne(query, selector interface{}) (UserRole, error) {
	ap := UserRole{}
	err := mongo.FindOne(db, roleCollection, query, selector, &ap)
	return ap, err
}
func (r UserRole) findAll(query, selector interface{}) (results []UserRole, err error) {
	results = []UserRole{}
	err = mongo.FindAll(db, roleCollection, query, selector, &results)
	return results, err
}

func (r UserRole) remove(selector interface{}) error {
	return mongo.Remove(db, roleCollection, selector)
}

func (r UserRole) removeAll(selector interface{}) error {
	return mongo.RemoveAll(db, roleCollection, selector)
}

func (r UserRole) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(db, roleCollection, query, selector)
}

func (r UserRole) findPage(page, limit int, query, selector interface{}, fields ...string) (results []UserRole, err error) {
	results = []UserRole{}
	err = mongo.FindPage(db, roleCollection, page, limit, query, selector, &results, fields...)
	return
}

func (r UserRole) Insert() error {
	if r.UserId == 0 || r.RoleId == 0 {
		return errors.New("userId & RoleId Needed")
	}
	r.Id, _ = mongo.GetIncrementId(roleCollection)
	if r.isExist(bson.M{"user_id": r.UserId, "role_id": r.RoleId}) {
		return fmt.Errorf("user role exist")
	}
	return r.insert(r)
}

func (r UserRole) Update() error {
	if r.isExist(bson.M{"user_id": r.UserId, "role_id": r.RoleId, "_id": bson.M{"$ne": r.Id}}) {
		return fmt.Errorf("角色已存在")
	}
	return r.update(bson.M{"_id": r.Id}, r)
}

func (r UserRole) Remove() error {
	if r.Id == 0 {
		return errors.New("id is 0")
	}
	return r.remove(bson.M{"_id": r.Id})
}
