package model_user

import (
	"encoding/json"
	"vue-admin/web_server/core/mongo"
	"vue-admin/web_server/model/mongo_index"
	"vue-admin/web_server/model/shareDB"

	"gopkg.in/mgo.v2/bson"
)

const (
	userCollection = mongo_index.CollectionUser
)

type User struct {
	Id         int64  `json:"id,omitempty" bson:"_id,omitempty"`
	Username   string `json:"username,omitempty" bson:"username,omitempty"` //用户名
	Password   string `json:"password,omitempty" bson:"password,omitempty"`
	Avatar     string `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Email      string `json:"email,omitempty" bson:"email,omitempty"`
	Phone      string `json:"phone,omitempty" bson:"phone,omitempty"`
	Gender     int    `json:"gender,omitempty" bson:"gender,omitempty"` // 1男 2女
	Name       string `json:"name,omitempty" bson:"name,omitempty"`     // 名字！
	Nick       string `json:"nick,omitempty" bson:"nick,omitempty"`     // 昵称
	Title      string `json:"title,omitempty" bson:"title,omitempty"`
	Status     int    `json:"status,omitempty" bson:"status,omitempty"` //1 激活，2锁定
	Note       string `json:"note,omitempty"  bson:"note,omitempty"`    //备注,
	CreateTime int64  `json:"create_time,omitempty"  bson:"create_time,omitempty"`
}

func (u User) ToJson() string {
	js, _ := json.Marshal(u)
	return string(js)
}

/*
	OperationModel
*/
func (u User) insert(docs ...interface{}) error {
	return mongo.Insert(shareDB.DBName(), userCollection, docs...)
}

func (u User) isExist(query interface{}) bool {
	return mongo.IsExist(shareDB.DBName(), userCollection, query)
}

func (u User) findOne(query, selector interface{}) (User, error) {
	us := User{}
	err := mongo.FindOne(shareDB.DBName(), userCollection, query, selector, &us)
	return us, err
}

func (u User) findAll(query, selector interface{}) (results []User, err error) {
	results = []User{}
	err = mongo.FindAll(shareDB.DBName(), userCollection, query, selector, &results)
	return results, err
}

func (u User) findPage(page, limit int, query, selector interface{}, fields ...string) (results []User, err error) {
	results = []User{}
	err = mongo.FindPage(shareDB.DBName(), userCollection, page, limit, query, selector, &results, fields...)
	return
}

//data := bson.M{"$set": bson.M{"age": 22}}
func (u User) update(selector, update interface{}) error {
	return mongo.Update(shareDB.DBName(), userCollection, selector, update, true)
}

func (u User) remove(selector interface{}) error {
	return mongo.Remove(shareDB.DBName(), userCollection, selector)
}

func (u User) removeAll(selector interface{}) error {
	return mongo.RemoveAll(shareDB.DBName(), userCollection, selector)
}

/*
	userModify
*/
func (u User) Insert() error {
	u.Id, _ = mongo.GetIncrementId(userCollection)
	u.Status = 1
	return u.insert(u)
}

func (u User) Update() error {
	selector := bson.M{"_id": u.Id}
	return u.update(selector, u)
}

func (u User) Remove() error {
	selector := bson.M{"_id": u.Id}
	return u.remove(selector)
}

func (u User) FindAll() ([]User, error) {
	return u.findAll(nil, bson.M{"password": 0})
}

func FindByUserId(userId int64) (u User, err error) {
	u = User{}
	u, err = u.findOne(bson.M{"_id": userId}, bson.M{"password": 0})
	return
}

func LoginUser(username, pass string) (user User, err error) {
	user = User{}
	user, err = user.findOne(bson.M{"username": username, "password": pass}, bson.M{"password": 0})
	return
}

func (u User) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(shareDB.DBName(), userCollection, query, selector)
}

func (u User) FindPageFilter(page, limit int, query, selector interface{}, fields ...string) ([]User, error) {
	return u.findPage(page, limit, query, selector, fields...)
}

func (u User) TotalCount(query, selector interface{}) (int, error) {
	return u.totalCount(query, selector)
}

func AddAdminUser() error {
	u := new(User)
	u.Username = "admin"
	u.Password = "flywithbug123"
	u.Email = "flywithbug@gmail.com"
	u.Title = "admin"
	u.Phone = "phone"
	u.Gender = 1
	return u.insert()
}