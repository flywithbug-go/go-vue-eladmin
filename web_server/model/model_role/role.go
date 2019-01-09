package model_role

import (
	"encoding/json"
	"fmt"
	"vue-admin/web_server/core/mongo"
	"vue-admin/web_server/model/shareDB"

	"gopkg.in/mgo.v2/bson"
)

const (
	roleCollection = "role"
)

//角色表，记录公司各种角色，比如：CEO 管理员，开发，开发经理，销售，销售主管，等
type Role struct {
	Id    int64 `json:"id,omitempty" bson:"_id,omitempty"`
	Pid   int64 `json:"pid"`
	Name  string
	Alias string
}

func (r Role) ToJson() string {
	js, _ := json.Marshal(r)
	return string(js)
}

func (r Role) isExist(query interface{}) bool {
	return mongo.IsExist(shareDB.DBName(), roleCollection, query)
}

func (r Role) insert(docs ...interface{}) error {
	return mongo.Insert(shareDB.DBName(), roleCollection, docs...)
}

func (r Role) update(selector, update interface{}) error {
	return mongo.Update(shareDB.DBName(), roleCollection, selector, update, true)
}

func (r Role) findOne(query, selector interface{}) (Role, error) {
	ap := Role{}
	err := mongo.FindOne(shareDB.DBName(), roleCollection, query, selector, &ap)
	return ap, err
}
func (r Role) findAll(query, selector interface{}) (results []Role, err error) {
	results = []Role{}
	err = mongo.FindAll(shareDB.DBName(), roleCollection, query, selector, &results)
	return results, err
}

func (r Role) remove(selector interface{}) error {
	return mongo.Remove(shareDB.DBName(), roleCollection, selector)
}

func (r Role) removeAll(selector interface{}) error {
	return mongo.RemoveAll(shareDB.DBName(), roleCollection, selector)
}

func (r Role) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(shareDB.DBName(), roleCollection, query, selector)
}

func (r Role) findPage(page, limit int, query, selector interface{}, fields ...string) (results []Role, err error) {
	results = []Role{}
	err = mongo.FindPage(shareDB.DBName(), roleCollection, page, limit, query, selector, &results, fields...)
	return
}

func (r Role) FindOne() (role Role, err error) {
	role, err = r.findOne(bson.M{"_id": r.Id}, nil)
	return
}
func (r Role) Insert() error {
	r.Id, _ = mongo.GetIncrementId(roleCollection)
	if r.isExist(bson.M{"name": r.Name}) {
		return fmt.Errorf("code exist")
	}
	if r.isExist(bson.M{"alias": r.Alias}) {
		return fmt.Errorf("alias exist")
	}
	return r.insert(r)
}

func (r Role) Update() error {
	return r.update(bson.M{"_id": r.Id}, r)
}

func (r Role) Remove() error {
	return r.remove(bson.M{"_id": r.Id})
}

func (r Role) TotalCount(query, selector interface{}) (int, error) {
	return r.totalCount(query, selector)
}
func (r Role) FindPageFilter(page, limit int, query, selector interface{}, fields ...string) ([]Role, error) {
	return r.findPage(page, limit, query, selector, fields...)
}
