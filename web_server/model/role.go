package model

import (
	"doc-manager/web_server/core/mongo"
	"encoding/json"
	"errors"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

const (
	roleCollection = "role"
)

//角色表，记录公司各种角色，比如：CEO 管理员，开发，开发经理，销售，销售主管，等
type Role struct {
	Id          int64  `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string `json:"name,omitempty"` //角色名称 CEO，CTO，主管，经理，程序员等。
	Code        string `json:"code"`           //角色编码
	DelFlag     bool   `json:"del_flag"`       //是否被删除
	Description string `json:"description"`
}

func (r Role) ToJson() string {
	js, _ := json.Marshal(r)
	return string(js)
}

func (r Role) isExist(query interface{}) bool {
	return mongo.IsExist(db, roleCollection, query)
}

func (r Role) insert(docs ...interface{}) error {
	return mongo.Insert(db, roleCollection, docs...)
}

func (r Role) update(selector, update interface{}) error {
	return mongo.Update(db, roleCollection, selector, update, true)
}

func (r Role) findOne(query, selector interface{}) (interface{}, error) {
	ap := Role{}
	err := mongo.FindOne(db, roleCollection, query, selector, &ap)
	return ap, err
}
func (r Role) findAll(query, selector interface{}) (results []Role, err error) {
	results = []Role{}
	err = mongo.FindAll(db, roleCollection, query, selector, &results)
	return results, err
}

func (r Role) remove(selector interface{}) error {
	return mongo.Remove(db, roleCollection, selector)
}

func (r Role) removeAll(selector interface{}) error {
	return mongo.RemoveAll(db, roleCollection, selector)
}

func (r Role) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(db, roleCollection, query, selector)
}

func (r Role) findPage(page, limit int, query, selector interface{}, fields ...string) (results []Role, err error) {
	results = []Role{}
	err = mongo.FindPage(db, roleCollection, page, limit, query, selector, &results, fields...)
	return
}

func (r Role) Insert() error {
	r.Id, _ = mongo.GetIncrementId(roleCollection)
	if r.isExist(bson.M{"code": r.Code}) {
		return fmt.Errorf("code exist")
	}

	if r.isExist(bson.M{"name": r.Name}) {
		return fmt.Errorf("name exist")
	}
	return r.insert(r)
}

func (r Role) Update() error {
	return r.update(bson.M{"_id": r.Id}, r)
}

func (r Role) Remove() error {
	if r.Id == 0 {
		return errors.New("id is 0")
	}
	return appC.remove(bson.M{"_id": r.Id})
}
