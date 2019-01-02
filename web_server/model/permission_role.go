package model

import (
	"doc-manager/web_server/core/mongo"
	"encoding/json"
	"errors"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

const (
	rolePermissionCollection = "role_permission"
)

type RolePermission struct {
	Id           int64 `json:"id" bson:"_id"`
	RoleId       int64
	PermissionId int64
}

func (r RolePermission) ToJson() string {
	js, _ := json.Marshal(r)
	return string(js)
}

func (r RolePermission) isExist(query interface{}) bool {
	return mongo.IsExist(db, rolePermissionCollection, query)
}

func (r RolePermission) insert(docs ...interface{}) error {
	return mongo.Insert(db, rolePermissionCollection, docs...)
}

func (r RolePermission) update(selector, update interface{}) error {
	return mongo.Update(db, rolePermissionCollection, selector, update, true)
}

func (r RolePermission) findOne(query, selector interface{}) (RolePermission, error) {
	ap := Role{}
	err := mongo.FindOne(db, rolePermissionCollection, query, selector, &ap)
	return ap, err
}
func (r RolePermission) findAll(query, selector interface{}) (results []RolePermission, err error) {
	results = []RolePermission{}
	err = mongo.FindAll(db, rolePermissionCollection, query, selector, &results)
	return results, err
}

func (r RolePermission) remove(selector interface{}) error {
	return mongo.Remove(db, rolePermissionCollection, selector)
}

func (r RolePermission) removeAll(selector interface{}) error {
	return mongo.RemoveAll(db, rolePermissionCollection, selector)
}

func (r RolePermission) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(db, rolePermissionCollection, query, selector)
}

func (r RolePermission) findPage(page, limit int, query, selector interface{}, fields ...string) (results []RolePermission, err error) {
	results = []RolePermission{}
	err = mongo.FindPage(db, rolePermissionCollection, page, limit, query, selector, &results, fields...)
	return
}

func (r RolePermission) Insert() error {
	if r.PermissionId == 0 || r.RoleId == 0 {
		return errors.New("permission_Id & RoleId Needed")
	}
	r.Id, _ = mongo.GetIncrementId(rolePermissionCollection)
	if r.isExist(bson.M{"permission_Id": r.PermissionId, "role_id": r.RoleId}) {
		return fmt.Errorf("permission_id role exist")
	}
	return r.insert(r)
}

func (r RolePermission) Update() error {
	if r.isExist(bson.M{"permission_id": r.PermissionId, "role_id": r.RoleId, "_id": bson.M{"$ne": r.Id}}) {
		return fmt.Errorf("角色已存在")
	}
	return r.update(bson.M{"_id": r.Id}, r)
}

func (r RolePermission) Remove() error {
	if r.Id == 0 {
		return errors.New("id is 0")
	}
	return appC.remove(bson.M{"_id": r.Id})
}
