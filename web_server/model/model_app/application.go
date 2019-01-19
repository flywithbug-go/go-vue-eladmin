package model_app

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"vue-admin/web_server/core/mongo"
	"vue-admin/web_server/model/a_mongo_index"
	"vue-admin/web_server/model/model_app_manager"
	"vue-admin/web_server/model/model_user"
	"vue-admin/web_server/model/shareDB"

	"gopkg.in/mgo.v2/bson"
)

const (
	appCollection = mongo_index.CollectionApp

	//ApplicationPermissionAll    = "APP_ALL"
	ApplicationPermissionSelect = "APP_SELECT"
	ApplicationPermissionCreate = "APP_CREATE"
	ApplicationPermissionEdit   = "APP_EDIT"
	ApplicationPermissionDelete = "APP_DELETE"
)

//修改规则，等级
// role 等级为1 的用户可以编辑
type Application struct {
	Id         int64             `json:"id,omitempty" bson:"_id,omitempty"`
	Name       string            `json:"name,omitempty" bson:"name,omitempty"`        //应用（组件）名称
	Desc       string            `json:"desc,omitempty" bson:"desc,omitempty"`        //项目描述
	CreateTime int64             `json:"time,omitempty" bson:"create_time,omitempty"` //创建时间
	Icon       string            `json:"icon,omitempty" bson:"icon,omitempty"`        //icon 地址
	Owner      string            `json:"owner,omitempty" bson:"owner,omitempty"`      //应用所有者
	OwnerId    int64             `json:"owner_id,omitempty" bson:"owner_id,omitempty"`
	BundleId   string            `json:"bundle_id,omitempty" bson:"bundle_id,omitempty"`
	Managers   []model_user.User `json:"manager,omitempty" bson:"manager,omitempty"`         //管理员
	ManagerIds []int64           `json:"manager_ids,omitempty" bson:"manager_ids,omitempty"` //管理员Id
}

func (a Application) ToJson() string {
	js, _ := json.Marshal(a)
	return string(js)
}

func (a Application) insert(docs ...interface{}) error {
	return mongo.Insert(shareDB.DocManagerDBName(), appCollection, docs...)
}

func (a Application) isExist(query interface{}) bool {
	return mongo.IsExist(shareDB.DocManagerDBName(), appCollection, query)
}

func (a Application) findOne(query, selector interface{}) (Application, error) {
	ap := Application{}
	err := mongo.FindOne(shareDB.DocManagerDBName(), appCollection, query, selector, &ap)
	return ap, err
}

func (a Application) findAll(query, selector interface{}) (results []Application, err error) {
	results = []Application{}
	err = mongo.FindAll(shareDB.DocManagerDBName(), appCollection, query, selector, &results)
	return results, err
}

func (a Application) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(shareDB.DocManagerDBName(), appCollection, query, selector)
}

func (a Application) findPage(page, limit int, query, selector interface{}, fields ...string) (results []Application, err error) {
	results = []Application{}
	err = mongo.FindPage(shareDB.DocManagerDBName(), appCollection, page, limit, query, selector, &results, fields...)
	return
}

func (a Application) update(selector, update interface{}) error {
	return mongo.Update(shareDB.DocManagerDBName(), appCollection, selector, update, true)
}

func (a Application) remove(selector interface{}) error {
	version := AppVersion{}
	if version.isExist(bson.M{"app_id": a.Id}) {
		return fmt.Errorf("app in use")
	}
	return mongo.Remove(shareDB.DocManagerDBName(), appCollection, selector)
}

func (a Application) removeAll(selector interface{}) error {
	return mongo.RemoveAll(shareDB.DocManagerDBName(), appCollection, selector)
}

func (a *Application) Insert() error {
	if a.BundleId == "" {
		return errors.New("bundleId must fill")
	}
	if a.Icon == "" {
		return errors.New("icon must fill")
	}
	if a.OwnerId == 0 {
		return errors.New("ownerId must fill")
	}
	if a.Name == "" {
		return errors.New("name must fill")
	}
	if len(a.Desc) < 10 {
		return errors.New("desc length must > 10")
	}

	if a.isExist(bson.M{"bundle_id": a.BundleId}) {
		return errors.New("bundle_id already exist")
	}
	if a.isExist(bson.M{"name": a.Name}) {
		return errors.New("name already exist")
	}
	a.Id, _ = mongo.GetIncrementId(shareDB.DocManagerDBName(), appCollection)
	a.CreateTime = time.Now().Unix() * 1000
	list := a.ManagerIds
	a.ManagerIds = nil
	err := a.insert(a)
	if err != nil {
		return err
	}
	a.ManagerIds = list
	a.updateAppManagers()
	return nil
}

func (a Application) updateAppManagers() error {
	if len(a.ManagerIds) == 0 {
		return nil
	}
	aM := model_app_manager.AppManager{}
	for _, userId := range a.ManagerIds {
		aM.UserId = userId
		aM.AppId = a.Id
		aM.Insert()
	}
	return nil
}

func (a Application) Update() error {
	a.BundleId = ""
	a.Owner = ""
	a.CreateTime = 0
	a.updateAppManagers()
	a.ManagerIds = nil
	return a.update(bson.M{"_id": a.Id}, a)
}

func (a Application) Remove() error {
	if a.Id == 0 {
		return errors.New("id is 0")
	}
	return a.remove(bson.M{"_id": a.Id})
}

func (a *Application) fetchManagers() error {
	user := model_user.User{}
	user.Id = a.OwnerId
	user, err := user.FindOne()
	if err != nil {
		return err
	}
	aM := model_app_manager.AppManager{}
	aMs, _ := aM.FindAll(bson.M{"app_id": a.Id}, nil)
	a.Managers = make([]model_user.User, 0)
	for _, item := range aMs {
		u := model_user.User{}
		u.Id = item.UserId
		u, err := u.FindOne()
		if err == nil {
			a.Managers = append(a.Managers, u)
		}
	}
	return nil
}

func (a Application) FindAll(query, selector interface{}) (apps []Application, err error) {
	apps, err = a.findAll(query, selector)
	for index := range apps {
		apps[index].fetchManagers()
	}
	return
}

//func FindPageApplications(page, limit int, fields ...string) (apps *[]Application, err error) {
//	return appC.findPage(page, limit, nil, nil, fields...)
//}

func (a Application) TotalCount(query, selector interface{}) (int, error) {
	return a.totalCount(query, selector)
}
func (a Application) FindPageFilter(page, limit int, query, selector interface{}, fields ...string) (apps []Application, err error) {
	apps, err = a.findPage(page, limit, query, selector, fields...)
	for index := range apps {
		apps[index].fetchManagers()
	}
	return
}
