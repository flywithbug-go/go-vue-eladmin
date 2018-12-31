package model

import (
	"doc-manager/web_server/common"
	"doc-manager/web_server/core/mongo"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
)

type appStatus int

var (
	appPlatformMap = map[string]string{"IOS": "iOS", "ANDROID": "Android", "H5": "H5", "SERVER": "Server"}
)

const (
	appVersionCollection              = "app_version"
	appStatusTypePrepare    appStatus = iota //准备中 待开发
	appStatusTypeDeveloping                  //开发中 待灰度
	appStatusTypeGray                        //灰度  待发布
	appStatusTypeRelease                     //已发布  已发布不能再更改
)

type AppVersion struct {
	Id            int64     `json:"id,omitempty" bson:"_id,omitempty"`
	AppId         int64     `json:"app_id,omitempty" bson:"app_id,omitempty"` //所属App DB Id
	Version       string    `json:"version,omitempty" bson:"version,omitempty"`
	ParentVersion string    `json:"parent_version,omitempty" bson:"parent_version,omitempty"`
	Platform      []string  `json:"platform,omitempty" bson:"platform,omitempty"`           //(iOS,Android,H5,Server)["iOS","Android","H5","Server"]
	Status        appStatus `json:"status,omitempty" bson:"status,omitempty"`               //状态    1(准备中) 2(开发中) 3(灰度) 4(已发布)
	ApprovalTime  int64     `json:"approval_time,omitempty" bson:"approval_time,omitempty"` //立项时间
	LockTime      int64     `json:"lock_time,omitempty" bson:"lock_time,omitempty"`         //锁版时间
	GrayTime      int64     `json:"gray_time,omitempty" bson:"gray_time,omitempty"`         //灰度时间
	CreateTime    int64     `json:"create_time,omitempty" bson:"create_time,omitempty"`     //添加时间
	AppStatus     string    `json:"app_status,omitempty" bson:"app_status,omitempty"`       //app状态
	ReleaseTime   int64     `json:"release_time,omitempty" bson:"release_time,omitempty"`
	Icon          string    `json:"icon,omitempty" bson:"icon,omitempty"`
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
	application, err := FindApplicationById(app.AppId)
	if err != nil {
		return fmt.Errorf("appID:%d,%s", app.AppId, err.Error())
	}
	app.Icon = application.Icon
	if app.isExist(bson.M{"version": app.Version, "app_id": app.AppId}) {
		return fmt.Errorf("version exist")
	}
	if len(app.ParentVersion) > 0 {
		if !app.isExist(bson.M{"version": app.ParentVersion, "app_id": app.AppId}) {
			return errors.New("parent_version not exist")
		}
	}
	for _, platform := range app.Platform {
		_, ok := appPlatformMap[strings.ToUpper(platform)]
		if !ok {
			return fmt.Errorf("platform must like (iOS,Android,H5,Server) ")
		}
	}
	app.Id, _ = mongo.GetIncrementId(appVersionCollection)
	app.CreateTime = time.Now().Unix()
	app.Status = appStatusTypePrepare
	app.AppStatus = makeStatusString(appStatusTypePrepare)
	compareState, err := common.VersionCompare(app.Version, app.ParentVersion)
	if err != nil {
		return err
	}
	if compareState != common.CompareVersionStateGreater {
		return errors.New("new Version must bigger than ParentVersion")
	}
	if len(app.ParentVersion) == 0 {
		app.ParentVersion = "-"
	}
	return appVC.insert(app)
}

func (app *AppVersion) Update() error {
	selector := bson.M{"_id": app.Id}
	app.ParentVersion = ""
	app.Version = ""
	app.CreateTime = 0
	app.AppStatus = makeStatusString(app.Status)
	return appVC.update(selector, app)
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

func TotalCountAppVersion(query, selector interface{}) (int, error) {
	return appVC.totalCount(query, selector)
}
func FindPageAppVersionFilter(page, limit int, query, selector interface{}, fields ...string) (apps *[]AppVersion, err error) {
	return appVC.findPage(page, limit, query, selector, fields...)
}

//func FindAppVersionByVersionId(appId string) (appV *AppVersion, err error) {
//	appV, err = appVC.findOne(bson.M{"app_id": appId}, nil)
//	appV.AppStatus = makeStatusString(appV.Status)
//	return
//}

func FindAppVersionById(Id int64) (appV *AppVersion, err error) {
	appV, err = appVC.findOne(bson.M{"_id": Id}, nil)
	appV.AppStatus = makeStatusString(appV.Status)
	return
}

func FindAppVersionByVersion(version string) (appV *AppVersion, err error) {
	appV, err = appVC.findOne(bson.M{"version": version}, nil)
	appV.AppStatus = makeStatusString(appV.Status)
	return
}

func FindAllAppVersion(query, selector interface{}) (results *[]AppVersion, err error) {
	return appVC.findAll(query, selector)
}
