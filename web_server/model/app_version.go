package model

import (
	"doc-manager/web_server/core/mongo"

	"github.com/globalsign/mgo/bson"
)

type appStatus int

const (
	appVersionCollection = "app_version"
)
const (
	appStatusTypePrepare    appStatus = iota //准备中
	appStatusTypeDeveloping                  //开发中
	appStatusTypeGray                        //灰度
	appStatusTypeRelease                     //已发布
)

type AppVersion struct {
	AppId         string    `json:"app_id,omitempty" bson:"app_id,omitempty"` //所属AppId
	Version       string    `json:"version,omitempty" bson:"version,omitempty"`
	ParentVersion string    `json:"parent_version,omitempty" bson:"parent_version,omitempty"`
	Platform      string    `json:"platform,omitempty" bson:"platform,omitempty"`           //(iOS,Android,H5,Server)["iOS","Android","H5","Server"]
	Status        appStatus `json:"status,omitempty" bson:"status,omitempty"`               //状态    0(准备中) 1(开发中) 2(灰度) 3(已发布)
	ApprovalTime  int       `json:"approval_time,omitempty" bson:"approval_time,omitempty"` //立项时间
	LockTime      int       `json:"lock_time,omitempty" bson:"lock_time,omitempty"`         //锁版时间
	GrayTime      int       `json:"gray_time,omitempty" bson:"gray_time,omitempty"`         //灰度时间
	CreateTime    int       `json:"create_time,omitempty" bson:"create_time,omitempty"`     //添加时间
	AppStatus     string    `json:"app_status,omitempty" bson:"app_status,omitempty"`       //app状态
}

var (
	appVC = AppVersion{}
)

func (app AppVersion) insert(docs ...interface{}) error {
	return mongo.Insert(db, appVersionCollection, docs...)
}

func (app AppVersion) isExist(query interface{}) bool {
	return mongo.IsExist(db, appVersionCollection, query)
}

func (app AppVersion) findOne(query, selector interface{}) (*AppVersion, error) {
	ap := new(AppVersion)
	err := mongo.FindOne(db, appVersionCollection, query, selector, ap)
	return ap, err
}

func (app AppVersion) findAll(query, selector interface{}) (results *[]AppVersion, err error) {
	results = new([]AppVersion)
	err = mongo.FindAll(db, appVersionCollection, query, selector, results)
	return results, err
}

func (app AppVersion) update(selector, update interface{}) error {
	return mongo.Update(db, appVersionCollection, selector, update, true)
}

func (app AppVersion) remove(selector interface{}) error {
	return mongo.Remove(db, appVersionCollection, selector)
}

func (app AppVersion) removeAll(selector interface{}) error {
	return mongo.RemoveAll(db, appVersionCollection, selector)
}

func (app AppVersion) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(db, appVersionCollection, query, selector)
}

func (app AppVersion) findPage(page, limit int, query, selector interface{}, fields ...string) (results *[]AppVersion, err error) {
	results = new([]AppVersion)
	err = mongo.FindPage(db, appVersionCollection, page, limit, query, selector, results, fields...)
	return
}

func (app *AppVersion) Insert() error {
	return appVC.insert(app)
}

func makeStatusString(status appStatus) string {
	statusString := "未知"
	switch status {
	case appStatusTypePrepare:
		statusString = "准备中"
	case appStatusTypeDeveloping:
		statusString = "开发中"
	case appStatusTypeGray:
		statusString = "灰度"
	case appStatusTypeRelease:
		statusString = "已发布"
	default:
		statusString = "待定"
	}
	return statusString
}

func FindPageAppVersionFilter(page, limit int, query, selector interface{}, fields ...string) (apps *[]AppVersion, err error) {
	return appVC.findPage(page, limit, query, selector, fields...)
}

func FindAppVersionByAppId(appId string) (appV *AppVersion, err error) {
	appV, err = appVC.findOne(bson.M{"app_id": appId}, nil)
	appV.AppStatus = makeStatusString(appV.Status)
	return
}
