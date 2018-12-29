package model

import "doc-manager/web_server/core/mongo"

type appStatus int

const (
	appVersionCollection = "app_version"

	appStatusTypePrepare    appStatus = iota //准备中
	appStatusTypeDeveloping                  //开发中
	appStatusTypeGray                        //灰度
	appStatusTypeRelease                     //已发布
)

type AppVersion struct {
	AppId         string    `json:"app_id"` //所属AppId
	Version       string    `json:"version"`
	ParentVersion string    `json:"parent_version"`
	Platform      string    `json:"platform"`      //(iOS,Android,H5,Server)["iOS","Android","H5","Server"]
	Status        appStatus `json:"status"`        //状态    0(准备中) 1(开发中) 2(灰度) 3(已发布)
	ApprovalTime  int       `json:"approval_time"` //立项时间
	LockTime      int       `json:"lock_time"`     //锁版时间
	GrayTime      int       `json:"gray_time"`     //灰度时间
	CreateTime    int       `json:"create_time"`   //添加时间
}

func (a *AppVersion) insert(docs ...interface{}) error {
	return mongo.Insert(db, appVersionCollection, docs...)
}

func (a *AppVersion) isExist(query interface{}) bool {
	return mongo.IsExist(db, appVersionCollection, query)
}

func (a *AppVersion) findOne(query, selector interface{}) (*AppVersion, error) {
	ap := new(AppVersion)
	err := mongo.FindOne(db, appVersionCollection, query, selector, ap)
	return ap, err
}

func (a *AppVersion) findAll(query, selector interface{}) (results *[]AppVersion, err error) {
	results = new([]AppVersion)
	err = mongo.FindAll(db, appVersionCollection, query, selector, results)
	return results, err
}

func (a *AppVersion) update(selector, update interface{}) error {
	return mongo.Update(db, appVersionCollection, selector, update, true)
}

func (a *AppVersion) remove(selector interface{}) error {
	return mongo.Remove(db, appVersionCollection, selector)
}

func (a *AppVersion) removeAll(selector interface{}) error {
	return mongo.RemoveAll(db, appVersionCollection, selector)
}

func (a *AppVersion) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(db, appVersionCollection, query, selector)
}

func (a *AppVersion) findPage(page, limit int, query, selector interface{}, fields ...string) (results *[]AppVersion, err error) {
	results = new([]AppVersion)
	err = mongo.FindPage(db, appVersionCollection, page, limit, query, selector, results, fields...)
	return
}
