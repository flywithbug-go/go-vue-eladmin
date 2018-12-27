package model

import (
	"doc-manager/web_server/core/mongo"
	"errors"

	"gopkg.in/mgo.v2/bson"
)

type roleState int

const (
	userCollection                = "user"
	roleStateTypeNormal roleState = iota //普通用户
	roleStateTypeAdmin                   //管理员
)

type User struct {
	Id       int64     `json:"_id" bson:"_id"`
	UserId   string    `json:"user_id" bson:"user_id"`
	Account  string    `json:"account"`
	Password string    `json:"password,omitempty"`
	Avatar   string    `json:"avatar,omitempty"`
	Email    string    `json:"email,omitempty"`
	Phone    string    `json:"phone,omitempty"`
	Sex      int       `json:"sex"` // 0保密，1男 2女
	RealName string    `json:"real_name" bson:"real_name"`
	Name     string    `json:"name,omitempty" bson:"name"` //昵称
	Title    string    `json:"title,omitempty"`
	Role     roleState `json:"role"`            //用于前端路由配置 1 管理员， 2 普通用户，
	Roles    []string  `json:"roles,omitempty"` //角色数组
	Status   int       `json:"status"`
}

var (
	userC = User{}
)

func (u User) Insert(docs ...interface{}) error {
	return mongo.Insert(db, userCollection, docs...)
}

func (u User) IsExist(query interface{}) bool {
	return mongo.IsExist(db, userCollection, query)
}

func (u User) FindOne(query, selector interface{}) (*User, error) {
	us := new(User)
	err := mongo.FindOne(db, userCollection, query, selector, us)
	return us, err
}

func (u User) FindAll(query, selector interface{}) (results *[]User, err error) {
	results = new([]User)
	err = mongo.FindAll(db, userCollection, query, selector, results)
	return results, err
}

func (u User) Update(selector, update interface{}) error {
	return mongo.Update(db, userCollection, selector, update)
}

func (u User) Remove(selector interface{}) error {
	return mongo.Remove(db, userCollection, selector)
}

func (u User) RemoveAll(selector interface{}) error {
	return mongo.RemoveAll(db, userCollection, selector)
}

func FindAllUsers() (*[]User, error) {
	return userC.FindAll(nil, bson.M{"password": 0})
}

func (u User) UserInsert() error {
	if mongo.IsExist(db, userCollection, bson.M{"account": u.Account}) {
		return errors.New("account 已存在")
	}
	if mongo.IsExist(db, userCollection, bson.M{"mail": u.Email}) {
		return errors.New("mail 已存在")
	}
	if u.UserId == "" {
		u.Id, _ = mongo.GetIncrementId("user")
		u.UserId = bson.NewObjectId().Hex()
	}
	return userC.Insert(u)
}

func LoginUser(account, pass string) (user *User, err error) {
	return userC.FindOne(bson.M{"account": account, "password": pass}, bson.M{"password": 0})

}

func FindByUserId(userId string) (u *User, err error) {
	return userC.FindOne(bson.M{"user_id": userId}, bson.M{"password": 0})
}

func MakeUserRoles(role roleState) []string {
	if role == roleStateTypeNormal {
		return []string{"normal"}
	} else if role == roleStateTypeAdmin {
		return []string{"admin"}
	} else {
		return []string{}
	}
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
	return u.Insert()
}
