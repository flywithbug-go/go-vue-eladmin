package model_role

import (
	"encoding/json"
	"fmt"
	"time"
	"vue-admin/web_server/core/mongo"
	"vue-admin/web_server/model/model_permission"
	"vue-admin/web_server/model/model_role_permission"
	"vue-admin/web_server/model/model_user_role"
	"vue-admin/web_server/model/mongo_index"
	"vue-admin/web_server/model/shareDB"

	"github.com/flywithbug/log4go"

	"gopkg.in/mgo.v2/bson"
)

const (
	roleCollection = mongo_index.CollectionRole
)

//角色表，记录公司各种角色，比如：CEO 管理员，开发，开发经理，销售，销售主管，等
type Role struct {
	Id          int64                         `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string                        `json:"name,omitempty"  bson:"name,omitempty"`
	Alias       string                        `json:"alias,omitempty"  bson:"alias,omitempty"`
	Note        string                        `json:"note,omitempty"  bson:"note,omitempty"`
	CreateTime  int64                         `json:"create_time,omitempty"  bson:"create_time,omitempty"`
	Permissions []model_permission.Permission `json:"permissions" bson:"permissions,omitempty"`
	Label       string                        `json:"label ,omitempty"  bson:"_ ,omitempty"`
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

func (r Role) Exist() bool {
	return r.isExist(bson.M{"_id": r.Id})
}

func (r Role) Insert() error {
	r.Id, _ = mongo.GetIncrementId(roleCollection)
	r.CreateTime = time.Now().Unix() * 1000
	list := r.Permissions
	r.Permissions = nil
	err := r.insert(r)
	if err != nil {
		return err
	}
	r.Permissions = list
	r.updateRolePermission()
	return nil
}

func (r Role) Update() error {
	if r.Id == 10000 {
		return fmt.Errorf("超级管理员不能编辑")
	}
	r.updateRolePermission()
	r.Permissions = nil
	selector := bson.M{"_id": r.Id}
	r.Id = 0
	return r.update(selector, r)
}

func (r *Role) updateRolePermission() {
	rp := model_role_permission.RolePermission{}
	rp.RemoveRoleId(r.Id)
	for _, per := range r.Permissions {
		if per.Exist() {
			rp.RoleId = r.Id
			rp.PermissionId = per.Id
			rp.Insert()
		}
	}
}

func (r Role) Remove() error {
	if r.Id == 10000 {
		return fmt.Errorf("超级管理员不能删除")
	}
	if r.checkInUse() {
		return fmt.Errorf("角色使用中，无法删除")
	}
	rp := model_role_permission.RolePermission{}
	rp.RemoveRoleId(r.Id)
	return r.remove(bson.M{"_id": r.Id})
}

func (r Role) FindOne() (role Role, err error) {
	return r.findOne(bson.M{"_id": r.Id}, nil)
}

func (r Role) TotalCount(query, selector interface{}) (int, error) {
	return r.totalCount(query, selector)
}

func (r Role) FindPageFilter(page, limit int, query, selector interface{}, fields ...string) ([]Role, error) {
	results, err := r.findPage(page, limit, query, selector, fields...)
	if err != nil {
		return nil, err
	}
	return results, err
}

func (r Role) FindOneTree() (role Role, err error) {
	role, err = r.findOne(bson.M{"_id": r.Id}, nil)
	if err != nil {
		return
	}
	list := []Role{r}
	makeTreeList(list, nil)
	return list[0], nil
}

func (r Role) FindPageTreeFilter(page, limit int, query, selector interface{}, fields ...string) ([]Role, error) {
	results, err := r.findPage(page, limit, query, selector, fields...)
	if err != nil {
		return nil, err
	}
	err = makeTreeList(results, selector)
	return results, err
}

func (r Role) FetchTreeList(selector interface{}) (results []Role, err error) {
	results, err = r.findAll(nil, selector)
	if err != nil {
		return
	}
	err = makeTreeList(results, selector)
	return
}

func makeTreeList(list []Role, selector interface{}) error {
	for index := range list {
		rp := model_role_permission.RolePermission{}
		results, _ := rp.FindAll(bson.M{"role_id": list[index].Id}, nil)
		list[index].Permissions = make([]model_permission.Permission, len(results))
		var per model_permission.Permission
		index1 := 0
		for _, item := range results {
			per.Id = item.PermissionId
			per, err := per.FindOne(selector)
			per.Label = per.Alias
			per.Alias = ""
			if err != nil {
				log4go.Info(err.Error())
			} else {
				list[index].Permissions[index1] = per
				index1++
			}
		}
		list[index].Permissions = list[index].Permissions[:index1]
	}
	return nil
}

func (r Role) checkInUse() bool {
	ur := model_user_role.UserRole{}
	if ur.Exist(bson.M{"role_id": r.Id}) {
		return true
	}
	return false
}
