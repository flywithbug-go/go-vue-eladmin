package model_app

import (
	"encoding/json"
	"errors"
	"time"
	"vue-admin/web_server/core/mongo"
	"vue-admin/web_server/model/shareDB"

	"gopkg.in/mgo.v2/bson"
)

const (
	appCollection = "application"
	role          = 1
)

//修改规则，等级
// role 等级为1 的用户可以编辑
type Application struct {
	Id         int64  `json:"id,omitempty" bson:"_id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`        //应用（组件）名称
	Desc       string `json:"desc,omitempty" bson:"desc,omitempty"`        //项目描述
	CreateTime int64  `json:"time,omitempty" bson:"create_time,omitempty"` //创建时间
	Icon       string `json:"icon,omitempty" bson:"icon,omitempty"`        //icon 地址
	Owner      string `json:"owner,omitempty" bson:"owner,omitempty"`      //负责人
	BundleId   string `json:"bundle_id,omitempty" bson:"bundle_id,omitempty"`
}

func (a Application) ToJson() string {
	js, _ := json.Marshal(a)
	return string(js)
}

func (a Application) insert(docs ...interface{}) error {
	return mongo.Insert(shareDB.DBName(), appCollection, docs...)
}

func (a Application) isExist(query interface{}) bool {
	return mongo.IsExist(shareDB.DBName(), appCollection, query)
}

func (a Application) findOne(query, selector interface{}) (Application, error) {
	ap := Application{}
	err := mongo.FindOne(shareDB.DBName(), appCollection, query, selector, &ap)
	return ap, err
}

func (a Application) findAll(query, selector interface{}) (results []Application, err error) {
	results = []Application{}
	err = mongo.FindAll(shareDB.DBName(), appCollection, query, selector, &results)
	return results, err
}

func (a Application) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(shareDB.DBName(), appCollection, query, selector)
}

func (a Application) findPage(page, limit int, query, selector interface{}, fields ...string) (results []Application, err error) {
	results = []Application{}
	err = mongo.FindPage(shareDB.DBName(), appCollection, page, limit, query, selector, &results, fields...)
	return
}

func (a Application) update(selector, update interface{}) error {
	return mongo.Update(shareDB.DBName(), appCollection, selector, update, true)
}

func (a Application) remove(selector interface{}) error {
	return mongo.Remove(shareDB.DBName(), appCollection, selector)
}

func (a Application) removeAll(selector interface{}) error {
	return mongo.RemoveAll(shareDB.DBName(), appCollection, selector)
}

func (a *Application) Insert() error {
	if a.BundleId == "" {
		return errors.New("bundleId must fill")
	}
	if a.Icon == "" {
		return errors.New("icon must fill")
	}
	if a.Owner == "" {
		return errors.New("owner must fill")
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
	a.Id, _ = mongo.GetIncrementId(appCollection)
	a.CreateTime = time.Now().Unix()
	return a.insert(a)
}

//func UpdateApplication(a *Application) error {
//	selector := bson.M{"_id": a.Id}
//	a.AppId = ""
//	a.BundleId = ""
//	a.Owner = ""
//	a.CreateTime = 0
//	return appC.update(selector, a)
//}

func (a Application) Update() error {
	selector := bson.M{}
	if a.Id > 0 {
		selector = bson.M{"_id": a.Id}
	} else {
		return errors.New("id & app_id is null")
	}
	a.BundleId = ""
	a.Owner = ""
	a.CreateTime = 0
	return a.update(selector, a)
}

func (a Application) Remove() error {
	if a.Id == 0 {
		return errors.New("id is 0")
	}
	return a.remove(bson.M{"_id": a.Id})
}

//func FindApplicationById(id int64) (Application, error) {
//	return a.findOne(bson.M{"_id": id}, nil)
//}
//
//func FindApplicationAppId(appId string) (Application, error) {
//	return a.findOne(bson.M{"app_id": appId}, nil)
//}
//func FindApplication(query, selector interface{}) (Application, error) {
//	return a.findOne(query, selector)
//}

func (a Application) FindAll(query, selector interface{}) (apps []Application, err error) {
	return a.findAll(query, selector)
}

//func FindPageApplications(page, limit int, fields ...string) (apps *[]Application, err error) {
//	return appC.findPage(page, limit, nil, nil, fields...)
//}

func (a Application) TotalCount(query, selector interface{}) (int, error) {
	return a.totalCount(query, selector)
}
func (a Application) FindPageFilter(page, limit int, query, selector interface{}, fields ...string) (apps []Application, err error) {
	return a.findPage(page, limit, query, selector, fields...)
}
