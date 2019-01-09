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
	permissionCollection = "permission"
)

const (
	PermissionTypeUndetermined typeStatus = iota
	PermissionTypeNone                    //无权限
	PermissionTypeCURD                    //增删改查
	PermissionTypeRU                      //改查
	PermissionTypeR                       //查看
)

//权限表 type
//1. Delete ReadWrite
//2. Read
//3. NoRight

type Permission struct {
	Id         int64      `json:"id,omitempty" bson:"_id,omitempty"`
	Type       typeStatus `json:"type,omitempty" bson:"type,omitempty"` //
	TypeStatus string     `json:"type_status,omitempty" bson:"type_status,omitempty"`
	Name       string     `json:"name,omitempty" bson:"name,omitempty"`         //
	Code       string     `json:"code,omitempty" bson:"code,omitempty"`         //
	DelFlag    bool       `json:"del_flag,omitempty" bson:"del_flag,omitempty"` //
	Note       string     `json:"note,omitempty" bson:"note,omitempty"`
}

func (p Permission) ToJson() string {
	js, _ := json.Marshal(p)
	return string(js)
}

func (p Permission) isExist(query interface{}) bool {
	return mongo.IsExist(shareDB.DBName(), permissionCollection, query)
}

func (p Permission) insert(docs ...interface{}) error {
	return mongo.Insert(shareDB.DBName(), permissionCollection, docs...)
}

func (p Permission) update(selector, update interface{}) error {
	return mongo.Update(shareDB.DBName(), permissionCollection, selector, update, true)
}

func (p Permission) findOne(query, selector interface{}) (Permission, error) {
	ap := Permission{}
	err := mongo.FindOne(shareDB.DBName(), permissionCollection, query, selector, &ap)
	return ap, err
}
func (p Permission) findAll(query, selector interface{}) (results []Permission, err error) {
	results = []Permission{}
	err = mongo.FindAll(shareDB.DBName(), permissionCollection, query, selector, &results)
	return results, err
}

func (p Permission) remove(selector interface{}) error {
	return mongo.Remove(shareDB.DBName(), permissionCollection, selector)
}

func (p Permission) removeAll(selector interface{}) error {
	return mongo.RemoveAll(shareDB.DBName(), permissionCollection, selector)
}

func (p Permission) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(shareDB.DBName(), permissionCollection, query, selector)
}

func (p Permission) findPage(page, limit int, query, selector interface{}, fields ...string) (results []Permission, err error) {
	results = []Permission{}
	err = mongo.FindPage(shareDB.DBName(), permissionCollection, page, limit, query, selector, &results, fields...)
	return
}

func (p Permission) Insert() error {
	p.Id, _ = mongo.GetIncrementId(permissionCollection)
	if p.isExist(bson.M{"code": p.Code}) {
		return fmt.Errorf("code exist")
	}
	if p.isExist(bson.M{"name": p.Name}) {
		return fmt.Errorf("name exist")
	}
	p.TypeStatus = makeTypeStatus(p.Type)
	return p.insert(p)
}

func (p Permission) FindOne() (Permission, error) {
	p, err := p.findOne(bson.M{"_id": p.Id}, nil)
	if err != nil {
		return p, err
	}
	p.TypeStatus = makeTypeStatus(p.Type)
	return p, err
}

func (p Permission) Update() error {
	if p.Id == 0 {
		return errors.New("id needed ")
	}
	return p.update(bson.M{"_id": p.Id}, p)
}

func (p Permission) Remove() error {
	if p.Id == 0 {
		return errors.New("id needed ")
	}
	return p.remove(bson.M{"_id": p.Id})
}

func makeTypeStatus(makeTypeStatus typeStatus) string {
	switch makeTypeStatus {
	case PermissionTypeUndetermined:
		return "无权限"
	case PermissionTypeNone:
		return "无权限"
	case PermissionTypeCURD:
		return "增删改查"
	case PermissionTypeRU:
		return "读写"
	case PermissionTypeR:
		return "只读"
	}
	return "未定义用户"
}

func (p Permission) TotalCount(query, selector interface{}) (int, error) {
	return p.totalCount(query, selector)
}
func (p Permission) FindPageFilter(page, limit int, query, selector interface{}, fields ...string) (apps []Permission, err error) {
	return p.findPage(page, limit, query, selector, fields...)
}
