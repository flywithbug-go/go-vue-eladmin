package model

import (
	"doc-manager/web_server/core/mongo"
)

const (
	appCollection = "application"
)

type Application struct {
	Id         int
	AppId      string `json:"app_id"`
	Name       string `json:"name"` //应用（组件）名称
	Desc       string //项目描述
	CreateTime int64  //创建时间
	Icon       string `json:"icon"` //icon 地址

	Owner string `json:"owner"` //负责人
}

var (
	app = Application{}
)

func (a Application) Insert(docs ...interface{}) error {
	return mongo.Insert(db, appCollection, docs...)
}

func (a Application) IsExist(query interface{}) bool {
	return mongo.IsExist(db, appCollection, query)
}

func (a Application) FindOne(query, selector interface{}) (*Application, error) {
	ap := new(Application)
	err := mongo.FindOne(db, appCollection, query, selector, ap)
	return ap, err
}

func (a Application) FindAll(query, selector interface{}) (results *[]Application, err error) {
	results = new([]Application)
	err = mongo.FindAll(db, appCollection, query, selector, results)
	return results, err
}

func (a Application) Update(selector, update interface{}) error {
	return mongo.Update(db, userCollection, selector, update, true)
}

func (a Application) Remove(selector interface{}) error {
	panic("implement me")
}

func (a Application) RemoveAll(selector interface{}) error {
	panic("implement me")
}
