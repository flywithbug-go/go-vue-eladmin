package model

import (
	"encoding/json"
	"errors"
	"vue-admin/web_server/core/mongo"

	"gopkg.in/mgo.v2/bson"
)

const (
	userCollection = "user"
)

type User struct {
	Id       int64  `json:"id,omitempty" bson:"_id,omitempty"`
	Account  string `json:"account,omitempty" bson:"account,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	Avatar   string `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Phone    string `json:"phone,omitempty" bson:"phone,omitempty"`
	Sex      int    `json:"sex,omitempty" bson:"sex,omitempty"`   // 0保密，1男 2女
	Name     string `json:"name,omitempty" bson:"name,omitempty"` // 名字！
	Nick     string `json:"nick,omitempty" bson:"nick,omitempty"` //昵称
	Title    string `json:"title,omitempty" bson:"title,omitempty"`
	Status   int    `json:"status,omitempty" bson:"status,omitempty"`
	Note     string `json:"note,omitempty"  bson:"note,omitempty"` //备注,
}

func (u User) ToJson() string {
	js, _ := json.Marshal(u)
	return string(js)
}

/*
	OperationModel
*/
func (u User) insert(docs ...interface{}) error {
	return mongo.Insert(db, userCollection, docs...)
}

func (u User) isExist(query interface{}) bool {
	return mongo.IsExist(db, userCollection, query)
}

func (u User) findOne(query, selector interface{}) (User, error) {
	us := User{}
	err := mongo.FindOne(db, userCollection, query, selector, &us)
	return us, err
}

func (u User) findAll(query, selector interface{}) (results []User, err error) {
	results = []User{}
	err = mongo.FindAll(db, userCollection, query, selector, &results)
	return results, err
}

func (u User) findPage(page, limit int, query, selector interface{}, fields ...string) (results []User, err error) {
	results = []User{}
	err = mongo.FindPage(db, appCollection, page, limit, query, selector, &results, fields...)
	return
}

//data := bson.M{"$set": bson.M{"age": 22}}
func (u User) update(selector, update interface{}) error {
	return mongo.Update(db, userCollection, selector, update, true)
}

func (u User) remove(selector interface{}) error {
	return mongo.Remove(db, userCollection, selector)
}

func (u User) removeAll(selector interface{}) error {
	return mongo.RemoveAll(db, userCollection, selector)
}

/*
	userModify
*/
func (u User) Insert() error {
	if u.isExist(bson.M{"account": u.Account}) {
		return errors.New("account 已存在")
	}
	if u.isExist(bson.M{"email": u.Email}) {
		return errors.New("email 已存在")
	}
	u.Id, _ = mongo.GetIncrementId(userCollection)
	return u.insert(u)
}

func (u User) Update() error {
	selector := bson.M{"_id": u.Id}
	u.Account = ""
	return u.update(selector, u)
}

func (u User) FindAll() ([]User, error) {
	return u.findAll(nil, bson.M{"password": 0})
}

func FindByUserId(userId int64) (u User, err error) {
	u = User{}
	u, err = u.findOne(bson.M{"_id": userId}, bson.M{"password": 0})
	return
}

func LoginUser(account, pass string) (user User, err error) {
	user = User{}
	user, err = user.findOne(bson.M{"account": account, "password": pass}, bson.M{"password": 0})
	return
}

func (u User) totalCount(query, selector interface{}) (int, error) {
	return mongo.TotalCount(db, appCollection, query, selector)
}

func (u User) FindPageFilter(page, limit int, query, selector interface{}, fields ...string) ([]User, error) {
	return u.findPage(page, limit, query, selector, fields...)
}

func (u User) TotalCount(query, selector interface{}) (int, error) {
	return u.totalCount(query, selector)
}

func AddAdminUser() error {
	u := new(User)
	u.Account = "admin"
	u.Password = "flywithbug123"
	u.Email = "flywithbug@gmail.com"
	u.Title = "admin"
	u.Phone = "phone"
	u.Sex = 1
	return u.insert()
}
