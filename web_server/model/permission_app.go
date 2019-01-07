package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"vue-admin/web_server/core/mongo"

	"gopkg.in/mgo.v2/bson"
)

const (
	appPermissionCollection = "app_permission"
)

type AppPermission struct {
	Id           int64 `json:"id,omitempty" bson:"_id,omitempty"`
	AppId        int64 `json:"app_id,omitempty" bson:"app_id,omitempty"`
	PermissionId int64 `json:"permission_id,omitempty" bson:"permission_id,omitempty"`
}

func (a AppPermission) ToJson() string {
	js, _ := json.Marshal(a)
	return string(js)
}

func (a AppPermission) isExist(query interface{}) bool {
	return mongo.IsExist(db, appPermissionCollection, query)
}

func (a AppPermission) insert(docs ...interface{}) error {
	return mongo.Insert(db, appPermissionCollection, docs...)
}

func (a AppPermission) update(selector, update interface{}) error {
	return mongo.Update(db, appPermissionCollection, selector, update, true)
}

func (a AppPermission) findOne(query, selector interface{}) (AppPermission, error) {
	ap := AppPermission{}
	err := mongo.FindOne(db, appPermissionCollection, query, selector, &ap)
	return ap, err
}
func (a AppPermission) findAll(query, selector interface{}) (results []AppPermission, err error) {
	results = []AppPermission{}
	err = mongo.FindAll(db, appPermissionCollection, query, selector, &results)
	return results, err
}

func (a AppPermission) remove(selector interface{}) error {
	return mongo.Remove(db, appPermissionCollection, selector)
}

func (a AppPermission) removeAll(selector interface{}) error {
	return mongo.RemoveAll(db, appPermissionCollection, selector)
}

func (a AppPermission) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(db, appPermissionCollection, query, selector)
}

func (a AppPermission) findPage(page, limit int, query, selector interface{}, fields ...string) (results []AppPermission, err error) {
	results = []AppPermission{}
	err = mongo.FindPage(db, appPermissionCollection, page, limit, query, selector, &results, fields...)
	return
}

func (a AppPermission) Insert() error {
	if a.PermissionId == 0 || a.AppId == 0 {
		return errors.New("permission_Id & AppId Needed")
	}
	a.Id, _ = mongo.GetIncrementId(appPermissionCollection)
	if a.isExist(bson.M{"permission_id": a.PermissionId, "app_id": a.AppId}) {
		return fmt.Errorf("permission_Id AppPermission exist")
	}
	return a.insert(a)
}

func (a AppPermission) Update() error {
	if a.isExist(bson.M{"permission_id": a.PermissionId, "app_id": a.AppId, "_id": bson.M{"$ne": a.Id}}) {
		return fmt.Errorf("角色已存在")
	}
	return a.update(bson.M{"_id": a.Id}, a)
}

func (a AppPermission) Remove() error {
	return a.remove(bson.M{"_id": a.Id})
}

func (a AppPermission) TotalCount(query, selector interface{}) (int, error) {
	return a.totalCount(query, selector)
}
func (a AppPermission) FindPageFilter(page, limit int, query, selector interface{}, fields ...string) (apps []AppPermission, err error) {
	return a.findPage(page, limit, query, selector, fields...)
}
