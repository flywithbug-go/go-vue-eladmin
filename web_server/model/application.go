package model

import (
	"doc-manager/web_server/core/mongo"
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	appCollection = "application"
)

type Application struct {
	Id         int64  `json:"id,omitempty" bson:"_id,omitempty"`
	AppId      string `json:"app_id,omitempty" bson:"app_id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`               //应用（组件）名称
	Desc       string `json:"desc,omitempty" bson:"desc,omitempty"`               //项目描述
	CreateTime int64  `json:"create_time,omitempty" bson:"create_time,omitempty"` //创建时间
	Icon       string `json:"icon,omitempty" bson:"icon,omitempty"`               //icon 地址
	Owner      string `json:"owner,omitempty" json:"owner,omitempty"`             //负责人 user_id，初始为创建人
	BundleId   string `json:"bundle_id,omitempty" bson:"bundle_id,omitempty"`
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

func (a Application) update(selector, update interface{}) error {
	return mongo.Update(db, userCollection, selector, update, true)
}

func (a Application) remove(selector interface{}) error {
	return mongo.Remove(db, appCollection, selector)
}

func (a Application) removeAll(selector interface{}) error {
	return mongo.RemoveAll(db, appCollection, selector)
}

func (a *Application) ApplicationInsert() error {
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

func (a *Application) ApplicationUpdate() error {
	query := bson.M{}
	if len(a.AppId) > 0 {
		query = bson.M{"app_id": a.AppId}
	} else if a.Id > 0 {
		query = bson.M{"_id": a.Id}
	} else {
		return errors.New("app_id or id not found")
	}
	a.BundleId = ""
	a.Name = ""
	return appC.update(query, a)
}

func FindALlApplications() (apps *[]Application, err error) {
	return appC.findAll(nil, nil)
}
