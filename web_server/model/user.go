package model

import (
	"doc-manager/web_server/core/mongo"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type  roleState int

const (
	userCollection = "user"
	roleStateTypeNormal roleState = iota  //普通用户
	roleStateTypeAdmin			//管理员
)


type User struct {
	Id       int64  `json:"_id" bson:"_id"`
	UserId   string `json:"user_id" bson:"user_id"`
	Account  string `json:"account"`
	Password string	`json:"password,omitempty"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email""`
	Phone    string `json:"phone"`
	Sex      int    `json:"sex"` // 0保密，1男 2女
	RealName string `json:"real_name" bson:"real_name"`
	Name     string `json:"name" bson:"name"`  //昵称
	Title    string `json:"title"`
	Role     roleState	`json:"role"`   //用于前端路由配置 1 管理员， 2 普通用户，
	Roles    []string `json:"roles"`   //角色数组
	Status   int	`json:"status"`
}



func FindAllUsers() ([]User, error) {
	var results []User
	err := mongo.FindAll(db, userCollection, nil, bson.M{"password":0}, &results)
	return results, err
}

func (u *User) Insert() error {
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
	return mongo.Insert(db, userCollection, u)
}

func (u *User) CheckLogin(account, pass string) (err error) {
	err = mongo.FindOne(db, userCollection, bson.M{"account": account, "password": pass}, nil, u)
	return err
}

func (u *User) Update() error {
	if u.UserId == "" {
		return errors.New("userId not equal id")
	}
	return mongo.Update(db, userCollection,
		bson.M{"_id": bson.ObjectIdHex(u.UserId)},
		bson.M{
			"avatar":u.Avatar,
			"role":u.Role,
			"roles":u.Roles,
			"title":u.Title,
			"status":u.Status,
			"name":u.Name,})
}

func (u User) Remove(id string) error {
	return mongo.Remove(db, userCollection, bson.M{"_id": bson.ObjectIdHex(id)})
}

func FindByUserId(userId string) (u *User, err error) {
	u = new(User)
	err = mongo.FindOne(db, userCollection, bson.M{"user_id": userId}, bson.M{"password":0}, &u)
	u.Roles = MakeUserRoles(u.Role)
	return
}

func MakeUserRoles(role roleState)[]string  {
	if role == roleStateTypeNormal {
		return []string{"normal"}
	}else if role == roleStateTypeAdmin {
		return []string{"admin"}
	}else {
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
