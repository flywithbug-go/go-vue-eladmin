package model_verify

import (
	"encoding/json"
	"fmt"
	"time"
	"vue-admin/web_server/core/mongo"
	"vue-admin/web_server/model/a_mongo_index"
	"vue-admin/web_server/model/shareDB"

	"math/rand"

	"gopkg.in/mgo.v2/bson"
)

const (
	verifyCollection = mongo_index.CollectionVerify
)

type Verify struct {
	Source     string `json:"source,omitempty" bson:"source,omitempty"`
	Code       string `json:"code,omitempty" bson:"code,omitempty"`
	Vld        int64  `json:"vld,omitempty" bson:"vld,omitempty"` //有效期
	CreateTime int64  `json:"create_time,omitempty" bson:"create_time,omitempty"`
	Marked     bool   `bson:"marked"`
}

func (v Verify) ToJson() string {
	js, _ := json.Marshal(v)
	return string(js)
}

func (v Verify) insert(docs ...interface{}) error {
	return mongo.Insert(shareDB.DBName(), verifyCollection, docs...)
}

func (v Verify) isExist(query interface{}) bool {
	return mongo.IsExist(shareDB.DBName(), verifyCollection, query)
}

func (v Verify) findOne(query, selector interface{}) (Verify, error) {
	ap := Verify{}
	err := mongo.FindOne(shareDB.DBName(), verifyCollection, query, selector, &ap)
	return ap, err
}

func (v Verify) findAll(query, selector interface{}) (results []Verify, err error) {
	results = []Verify{}
	err = mongo.FindAll(shareDB.DBName(), verifyCollection, query, selector, &results)
	return results, err
}

func (v Verify) update(selector, update interface{}) error {
	return mongo.Update(shareDB.DBName(), verifyCollection, selector, update, true)
}

func (v Verify) remove(selector interface{}) error {
	return mongo.Remove(shareDB.DBName(), verifyCollection, selector)
}

func (v Verify) removeAll(selector interface{}) error {
	return mongo.RemoveAll(shareDB.DBName(), verifyCollection, selector)
}

func (v Verify) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(shareDB.DBName(), verifyCollection, query, selector)
}

func (v Verify) findPage(page, limit int, query, selector interface{}, fields ...string) (results []Verify, err error) {
	results = []Verify{}
	err = mongo.FindPage(shareDB.DBName(), verifyCollection, page, limit, query, selector, &results, fields...)
	return
}

func (v Verify) Insert() error {
	v.CreateTime = time.Now().Unix()
	v.Marked = false
	return v.insert(v)
}

func GeneralVerifyData(source string) (string, error) {
	var verify Verify
	verify.Code = verify.generalVCode(source)
	verify.Source = source
	verify.Marked = false
	verify.Vld = time.Now().Unix() + 300
	err := verify.Insert()
	return verify.Code, err
}

func (v Verify) generalVCode(source string) string {
	rand.Int()
	vCode := createCaptcha()
	if v.isExist(bson.M{"marked": false, "code": vCode, "source": source, "vld": bson.M{"$gte": time.Now().Unix()}}) {
		vCode = v.generalVCode(source)
	}
	return vCode
}

func CheckVerify(source, code string) bool {
	var verify Verify
	if verify.isExist(bson.M{"marked": false, "code": code, "source": source, "vld": bson.M{"$gte": time.Now().Unix()}}) {
		updateMarked(source, code)
		return true
	}
	return false
}

func updateMarked(source, code string) {
	var verify Verify
	verify.update(bson.M{"code": code, "source": source, "marked": false}, bson.M{"marked": true})
}

func createCaptcha() string {
	return fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}
