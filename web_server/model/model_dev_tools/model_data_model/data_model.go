package model_data_model

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
	"vue-admin/web_server/core/mongo"
	"vue-admin/web_server/model/a_mongo_index"
	"vue-admin/web_server/model/shareDB"

	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

type typeStatus int

const (
	modelAttributeTypeUndefine typeStatus = iota //待定
	//基础类型
	modelAttributeTypeInt    //Int类型
	modelAttributeTypeBool   //布尔类型
	modelAttributeTypeString //String类型
	modelAttributeTypeList   //数组 （基础类型或者模型）
	modelAttributeTypeObject //模型
)

const (
	dataModelCollection = mongo_index.CollectionDataModel
)

type Attribute struct {
	Type      typeStatus `json:"type,omitempty" bson:"type,omitempty"` //int string list bool
	Name      string     `json:"name,omitempty" bson:"name,omitempty"`
	ModelName string     `json:"model_name,omitempty" bson:"model_name,omitempty"`
	ModelId   int64      `json:"model_id,omitempty" bson:"model_id,omitempty"`
}

type DataModel struct {
	Id         int64       `json:"id,omitempty" bson:"_id,omitempty"`
	Name       string      `json:"name,omitempty" bson:"name,omitempty"`
	Desc       string      `json:"desc,omitempty" bson:"desc,omitempty"`
	CreateTime int64       `json:"create_time,omitempty" bson:"create_time,omitempty"`
	Attributes []Attribute `json:"attributes,omitempty" bson:"attributes,omitempty"` //模型的属性表
}

func (d DataModel) ToJson() string {
	js, _ := json.Marshal(d)
	return string(js)
}

func (d DataModel) isExist(query interface{}) bool {
	return mongo.IsExist(shareDB.DocManagerDBName(), dataModelCollection, query)
}

func (d DataModel) insert(docs ...interface{}) error {
	return mongo.Insert(shareDB.DocManagerDBName(), dataModelCollection, docs...)
}

func (d DataModel) update(selector, update interface{}) error {
	return mongo.Update(shareDB.DocManagerDBName(), dataModelCollection, selector, update, true)
}

func (d DataModel) findOne(query, selector interface{}) (DataModel, error) {
	ap := DataModel{}
	err := mongo.FindOne(shareDB.DocManagerDBName(), dataModelCollection, query, selector, &ap)
	return ap, err
}
func (d DataModel) findAll(query, selector interface{}) (results []DataModel, err error) {
	results = []DataModel{}
	err = mongo.FindAll(shareDB.DocManagerDBName(), dataModelCollection, query, selector, &results)
	return results, err
}

func (d DataModel) remove(selector interface{}) error {
	return mongo.Remove(shareDB.DocManagerDBName(), dataModelCollection, selector)
}

func (d DataModel) removeAll(selector interface{}) error {
	return mongo.RemoveAll(shareDB.DocManagerDBName(), dataModelCollection, selector)
}

func (d DataModel) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(shareDB.DocManagerDBName(), dataModelCollection, query, selector)
}

func (d DataModel) findPage(page, limit int, query, selector interface{}, fields ...string) (results []DataModel, err error) {
	results = []DataModel{}
	err = mongo.FindPage(shareDB.DocManagerDBName(), dataModelCollection, page, limit, query, selector, &results, fields...)
	return
}

func (d DataModel) pipeAll(pipeline, result interface{}, allowDiskUse bool) error {
	return mongo.PipeAll(shareDB.DocManagerDBName(), dataModelCollection, pipeline, result, allowDiskUse)
}

func (d DataModel) pipeOne(pipeline, result interface{}, allowDiskUse bool) error {
	return mongo.PipeOne(shareDB.DocManagerDBName(), dataModelCollection, pipeline, result, allowDiskUse)
}

func (d DataModel) explain(pipeline, result interface{}) (results []DataModel, err error) {
	err = mongo.Explain(shareDB.DocManagerDBName(), dataModelCollection, pipeline, result)
	return
}

func (d DataModel) Update() error {
	return d.update(bson.M{"_id": d.Id}, d)
}

func (d DataModel) AddAttribute(a Attribute) error {
	if d.isExistAttribute(a) {
		return fmt.Errorf("attribute is exist")
	}
	update := bson.M{"$addToSet": bson.M{"attributes": a}}
	change := mgo.Change{
		Update: update,
	}
	query := bson.M{"_id": d.Id}
	ms, c := mongo.Collection(shareDB.DocManagerDBName(), dataModelCollection)
	defer ms.Close()
	_, err := c.Find(query).Apply(change, nil)
	return err
}

func (d DataModel) isExistAttribute(a Attribute) bool {
	selector := bson.M{"_id": d.Id, "attributes.name": a.Name}
	return d.isExist(selector)
}

func (d DataModel) Insert() (id int64, err error) {
	id, err = mongo.GetIncrementId(shareDB.DocManagerDBName(), dataModelCollection)
	if err != nil {
		return -1, err
	}
	d.CreateTime = time.Now().Unix()
	d.Id = id
	return id, d.insert(d)
}

func (d DataModel) TotalCount(query, selector interface{}) (int, error) {
	return d.totalCount(query, selector)
}
func (d DataModel) FindPageFilter(page, limit int, query, selector interface{}, fields ...string) ([]DataModel, error) {
	return d.findPage(page, limit, query, selector, fields...)
}

func duplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}
