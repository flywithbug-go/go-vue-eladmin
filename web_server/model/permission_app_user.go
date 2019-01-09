package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"vue-admin/web_server/core/mongo"
	"vue-admin/web_server/model/shareDB"

	"gopkg.in/mgo.v2/bson"
)

const (
	userAppPermissionCollection = "user_app_permission"
)

//用户关联App的管理权限
type UserAppPermission struct {
	Id              int64 `json:"id,omitempty" bson:"_id,omitempty"`
	UserId          int64 `json:"user_id,omitempty" bson:"user_id,omitempty"`
	AppPermissionId int64 `json:"app_permission_id,omitempty" bson:"app_permission_id,omitempty"`
}

func (r UserAppPermission) ToJson() string {
	js, _ := json.Marshal(r)
	return string(js)
}

func (r UserAppPermission) isExist(query interface{}) bool {
	return mongo.IsExist(shareDB.DBName(), userAppPermissionCollection, query)
}

func (r UserAppPermission) insert(docs ...interface{}) error {
	return mongo.Insert(shareDB.DBName(), userAppPermissionCollection, docs...)
}

func (r UserAppPermission) update(selector, update interface{}) error {
	return mongo.Update(shareDB.DBName(), userAppPermissionCollection, selector, update, true)
}

func (r UserAppPermission) findOne(query, selector interface{}) (UserAppPermission, error) {
	ap := UserAppPermission{}
	err := mongo.FindOne(shareDB.DBName(), userAppPermissionCollection, query, selector, &ap)
	return ap, err
}
func (r UserAppPermission) findAll(query, selector interface{}) (results []UserAppPermission, err error) {
	results = []UserAppPermission{}
	err = mongo.FindAll(shareDB.DBName(), userAppPermissionCollection, query, selector, &results)
	return results, err
}

func (r UserAppPermission) remove(selector interface{}) error {
	return mongo.Remove(shareDB.DBName(), userAppPermissionCollection, selector)
}

func (r UserAppPermission) removeAll(selector interface{}) error {
	return mongo.RemoveAll(shareDB.DBName(), userAppPermissionCollection, selector)
}

func (r UserAppPermission) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(shareDB.DBName(), userAppPermissionCollection, query, selector)
}

func (r UserAppPermission) findPage(page, limit int, query, selector interface{}, fields ...string) (results []UserAppPermission, err error) {
	results = []UserAppPermission{}
	err = mongo.FindPage(shareDB.DBName(), userAppPermissionCollection, page, limit, query, selector, &results, fields...)
	return
}

func (r UserAppPermission) Insert() error {
	if r.AppPermissionId == 0 || r.UserId == 0 {
		return errors.New("app_permission_id & UserId Needed")
	}
	r.Id, _ = mongo.GetIncrementId(userAppPermissionCollection)
	if r.isExist(bson.M{"app_permission_id": r.AppPermissionId, "user_id": r.UserId}) {
		return fmt.Errorf("permission_id role exist")
	}
	return r.insert(r)
}

func (r UserAppPermission) Update() error {
	if r.isExist(bson.M{"app_permission_id": r.AppPermissionId, "user_id": r.UserId, "_id": bson.M{"$ne": r.Id}}) {
		return fmt.Errorf("角色已存在")
	}
	return r.update(bson.M{"_id": r.Id}, r)
}

func (r UserAppPermission) Remove() error {
	if r.Id == 0 {
		return errors.New("id is 0")
	}
	return r.remove(bson.M{"_id": r.Id})
}

func (r UserAppPermission) TotalCount(query, selector interface{}) (int, error) {
	return r.totalCount(query, selector)
}
func (r UserAppPermission) FindPageFilter(page, limit int, query, selector interface{}, fields ...string) ([]UserAppPermission, error) {
	return r.findPage(page, limit, query, selector, fields...)
}
