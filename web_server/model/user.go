package model

import (
	"doc-manager/web_server/core/mongo"
	"encoding/json"
	"errors"

	"gopkg.in/mgo.v2/bson"
)

const (
	userCollection         = "user"
	roleStateTypeForbidden = -1 //禁用
	roleStateTypeAdmin     = 1  //管理员
	roleStateTypeDeveloper = 2  //开发者
	roleStateTypeNormal    = 3  //普通用户
	roleStateTypeRootAdmin = 11 //最高管理员
)

type User struct {
	Id       int64  `json:"id,omitempty" bson:"_id,omitempty"`
	Account  string `json:"account,omitempty" bson:"account,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	Avatar   string `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Phone    string `json:"phone,omitempty" bson:"phone,omitempty"`
	Sex      int    `json:"sex,omitempty" bson:"sex,omitempty"` // 0保密，1男 2女
	RealName string `json:"real_name,omitempty" bson:"real_name,omitempty"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"` //昵称
	Title    string `json:"title,omitempty" bson:"title,omitempty"`
	Status   int    `json:"status,omitempty" bson:"status,omitempty"`
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

func makeUserRoles(role int) string {
	switch role {
	case roleStateTypeRootAdmin:
		return "最高管理员"
	case roleStateTypeAdmin:
		return "管理员"
	case roleStateTypeDeveloper:
		return "开发者"
	case roleStateTypeNormal:
		return "普通用户"
	case roleStateTypeForbidden:
		return "被禁用户"
	}
	return "未定义用户"
}

func AddAdminUser() error {
	u := new(User)

	u.Account = "admin"
	u.Password = "flywithbug123"
	u.Email = "flywithbug@gmail.com"
	u.Title = "admin"
	u.Phone = "phone"
	u.RealName = "Jack"
	u.Sex = 1
	return u.insert()
}
