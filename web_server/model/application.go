package model

import (
	"doc-manager/web_server/core/mongo"
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	appCollection = "application"
	role          = 1
)

//修改规则，等级
// role 等级为1 的用户可以编辑
type Application struct {
	Id         int64    `json:"id,omitempty" bson:"_id,omitempty"`
	AppId      string   `json:"app_id,omitempty" bson:"app_id,omitempty"`
	Name       string   `json:"name,omitempty" bson:"name,omitempty"`        //应用（组件）名称
	Desc       string   `json:"desc,omitempty" bson:"desc,omitempty"`        //项目描述
	CreateTime int64    `json:"time,omitempty" bson:"create_time,omitempty"` //创建时间
	Icon       string   `json:"icon,omitempty" bson:"icon,omitempty"`        //icon 地址
	Owner      string   `json:"owner,omitempty" bson:"owner,omitempty"`      //负责人
	BundleId   string   `json:"bundle_id,omitempty" bson:"bundle_id,omitempty"`
	Manager    []string `json:"manager,omitempty" bson:"manager,omitempty"` //管理者 account 数组
	Editable   bool     //当前用户是否可以编辑
}

var (
	appC = Application{}
)

func (a Application) insert(docs ...interface{}) error {
	return mongo.Insert(db, appCollection, docs...)
}

func (a Application) isExist(query interface{}) bool {
	return mongo.IsExist(db, appCollection, query)
}

func (a Application) findOne(query, selector interface{}) (*Application, error) {
	ap := new(Application)
	err := mongo.FindOne(db, appCollection, query, selector, ap)
	return ap, err
}

func (a Application) findAll(query, selector interface{}) (results *[]Application, err error) {
	results = new([]Application)
	err = mongo.FindAll(db, appCollection, query, selector, results)
	return results, err
}

func (a Application) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(db, appCollection, query, selector)
}

func (a Application) findPage(page, limit int, query, selector interface{}, fields ...string) (results *[]Application, err error) {
	results = new([]Application)
	err = mongo.FindPage(db, appCollection, page, limit, query, selector, results, fields...)
	return
}

func (a Application) update(selector, update interface{}) error {
	return mongo.Update(db, appCollection, selector, update, true)
}

func (a Application) remove(selector interface{}) error {
	return mongo.Remove(db, appCollection, selector)
}

func (a Application) removeAll(selector interface{}) error {
	return mongo.RemoveAll(db, appCollection, selector)
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

	if appC.isExist(bson.M{"bundle_id": a.BundleId}) {
		return errors.New("bundle_id already exist")
	}
	if appC.isExist(bson.M{"name": a.Name}) {
		return errors.New("name already exist")
	}
	a.AppId = bson.NewObjectId().Hex()
	a.Id, _ = mongo.GetIncrementId(appCollection)
	a.CreateTime = time.Now().Unix()
	return appC.insert(a)
}

//func UpdateApplication(a *Application) error {
//	selector := bson.M{"_id": a.Id}
//	a.AppId = ""
//	a.BundleId = ""
//	a.Owner = ""
//	a.CreateTime = 0
//	return appC.update(selector, a)
//}

func (a *Application) Update() error {
	selector := bson.M{}
	if a.Id > 0 {
		selector = bson.M{"_id": a.Id}
	} else if len(a.AppId) > 0 {
		selector = bson.M{"app_id": a.AppId}
	} else {
		return errors.New("id & app_id is null")
	}
	a.AppId = ""
	a.BundleId = ""
	a.Owner = ""
	a.CreateTime = 0
	return appC.update(selector, a)
}

func FindApplicationById(id int64) (*Application, error) {
	return appC.findOne(bson.M{"_id": id}, nil)
}

func FindApplicationAppId(appId string) (*Application, error) {
	return appC.findOne(bson.M{"app_id": appId}, nil)
}
func FindApplication(query, selector interface{}) (*Application, error) {
	return appC.findOne(query, selector)
}

//func FindALlApplications() (apps *[]Application, err error) {
//	return appC.findAll(nil, nil)
//}

//func FindPageApplications(page, limit int, fields ...string) (apps *[]Application, err error) {
//	return appC.findPage(page, limit, nil, nil, fields...)
//}

func TotalCountApplication(query, selector interface{}) (int, error) {
	return appC.totalCount(query, selector)
}
func FindPageApplicationsFilter(page, limit int, query, selector interface{}, fields ...string) (apps *[]Application, err error) {
	return appC.findPage(page, limit, query, selector, fields...)
}
