package model

import (
	"doc-manager/web_server/core/mongo"
	"encoding/json"
)

const (
	permissionCollection = "role"
)

//权限表 type
//1. Delete ReadWrite
//2. Read
//3. NoRight

type Permission struct {
	Id          int64  `json:"id,omitempty" bson:"_id,omitempty"`
	Type        int    `json:"type"`           //
	Name        string `json:"name,omitempty"` //
	Code        string `json:"code"`           //
	DelFlag     bool   `json:"del_flag"`       //
	Description string `json:"description"`
}

func (p Permission) ToJson() string {
	js, _ := json.Marshal(p)
	return string(js)
}

func (p Permission) isExist(query interface{}) bool {
	return mongo.IsExist(db, permissionCollection, query)
}

func (p Permission) insert(docs ...interface{}) error {
	return mongo.Insert(db, permissionCollection, docs...)
}

func (p Permission) update(selector, update interface{}) error {
	return mongo.Update(db, permissionCollection, selector, update, true)
}

func (p Permission) findOne(query, selector interface{}) (interface{}, error) {
	ap := Permission{}
	err := mongo.FindOne(db, permissionCollection, query, selector, &ap)
	return ap, err
}
func (p Permission) findAll(query, selector interface{}) (results []Permission, err error) {
	results = []Permission{}
	err = mongo.FindAll(db, permissionCollection, query, selector, &results)
	return results, err
}

func (p Permission) remove(selector interface{}) error {
	return mongo.Remove(db, permissionCollection, selector)
}

func (p Permission) removeAll(selector interface{}) error {
	return mongo.RemoveAll(db, permissionCollection, selector)
}

func (p Permission) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(db, permissionCollection, query, selector)
}

func (p Permission) findPage(page, limit int, query, selector interface{}, fields ...string) (results []Permission, err error) {
	results = []Permission{}
	err = mongo.FindPage(db, permissionCollection, page, limit, query, selector, &results, fields...)
	return
}

func (p Permission) Insert() error {
	panic("implement me")
}

func (p Permission) Update() error {
	panic("implement me")
}
