package model

import (
	"errors"

	"gopkg.in/mgo.v2/bson"

	"doc-manager/web_server/core/mongo"
)

type roleState int

const (
	userCollection                = "user"
	roleStateTypeNormal roleState = iota //普通用户
	roleStateTypeAdmin                   //管理员
)

type User struct {
	Id       int64     `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId   string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Account  string    `json:"account,omitempty" bson:"account,omitempty"`
	Password string    `json:"password,omitempty" bson:"password,omitempty"`
	Avatar   string    `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Email    string    `json:"email,omitempty" bson:"email,omitempty"`
	Phone    string    `json:"phone,omitempty" bson:"phone,omitempty"`
	Sex      int       `json:"sex,omitempty" bson:"sex,omitempty"` // 0保密，1男 2女
	RealName string    `json:"real_name,omitempty" bson:"real_name,omitempty"`
	Name     string    `json:"name,omitempty" bson:"name,omitempty"` //昵称
	Title    string    `json:"title,omitempty" bson:"title,omitempty"`
	Role     roleState `json:"role,omitempty" bson:"role,omitempty"`   //用于前端路由配置 1 管理员， 2 普通用户，
	Roles    []string  `json:"roles,omitempty" bson:"roles,omitempty"` //角色数组
	Status   int       `json:"status,omitempty" bson:"status,omitempty"`
}

var (
	userC = new(User)
)

/*
	OperationModel
*/
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

//data := bson.M{"$set": bson.M{"age": 22}}
func (u User) Update(selector, update interface{}) error {
	return mongo.Update(db, userCollection, selector, update, true)
}

func (u User) Remove(selector interface{}) error {
	return mongo.Remove(db, userCollection, selector)
}

func (u User) RemoveAll(selector interface{}) error {
	return mongo.RemoveAll(db, userCollection, selector)
}

/*
	userModify
*/
func (u *User) UserInsert() error {
	if u.IsExist(bson.M{"account": u.Account}) {
		return errors.New("account 已存在")
	}
	if u.IsExist(bson.M{"email": u.Email}) {
		return errors.New("email 已存在")
	}
	if u.UserId == "" {
		u.Id, _ = mongo.GetIncrementId("user")
		u.UserId = bson.NewObjectId().Hex()
	}
	if u.Role == 0 {
		u.Role = 2
		u.Roles = makeUserRoles(u.Role)
	}
	return userC.Insert(u)
}

func UpdateUserInfo(u *User) error {
	return userC.Update(bson.M{"user_id": u.UserId}, u)
}

func FindAllUsers() (*[]User, error) {
	return userC.FindAll(nil, bson.M{"password": 0})
}

func FindByUserId(userId string) (u *User, err error) {
	u, err = userC.FindOne(bson.M{"user_id": userId}, bson.M{"password": 0})
	u.Roles = makeUserRoles(u.Role)
	return
}

func LoginUser(account, pass string) (user *User, err error) {
	user, err = userC.FindOne(bson.M{"account": account, "password": pass}, bson.M{"password": 0})
	user.Roles = makeUserRoles(user.Role)
	return
}

func makeUserRoles(role roleState) []string {
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
