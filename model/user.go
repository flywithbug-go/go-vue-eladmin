package model

import (
	"errors"
	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type ModelOperation interface {
	FindAll()( error)
	Insert()error
	Update(id string)error
	Remove(id string)error
	FindById(id string)(interface{}, error)
}

const (
	db         = "doc-manager"
	userCollection = "user"
	userRoleCollection = "user_role"
	appCollection = "application"

)


type User struct {
	Id			bson.ObjectId 	`json:"_id" bson:"_id"`
	UserId 		string			`json:"user_id" bson:"user_id"`
	Account 	string  		//`json:"account"`  //登录账号（唯一）
	Password 	string
	Mail 		string			//
	Phone 		string
	Sex 		string
	RealName	string			`json:"real_name" bson:"real_name"`
	Title 		string  //职位
}

func (u User) FindAll() ([]User,error) {
	var results []User
	err := FindAll(db,userCollection,nil,nil,&results)
	return results,err
}

func (u User) Insert() error {
	if IsExist(db,userCollection,bson.M{"account":u.Account}) {
		return errors.New("account 已存在")
	}
	if IsExist(db,userCollection,bson.M{"mail":u.Mail}) {
		return errors.New("mail 已存在")
	}
	if u.UserId == "" {
		u.Id = bson.NewObjectId()
		u.UserId = u.Id.Hex()
	}
	return Insert(db,userCollection,u)
}

func (u User) Update(id string) error {
	if !strings.EqualFold(u.UserId,id) {
		return errors.New("userId not equal id")
	}
	return Update(db,userCollection,bson.M{"_id":bson.ObjectIdHex(id)},u)
}

func (u User) Remove(id string) error {
	return Remove(db,userCollection,bson.M{"_id":bson.ObjectIdHex(id)})
}

func (u User) FindById(id string) (*User, error) {
	err := FindOne(db,userCollection,bson.M{"_id":bson.ObjectIdHex(id)},nil,&u)
	return &u,err
}
func (u User)UserLogin(account, password string)(*User,error)  {
	err := FindOne(db, userCollection,bson.M{"account":account,"password":password},nil,&u)
	return &u,err
}

func (u User)FindByQuery(query mgo.Query)(*User, error)  {
	err := FindOne(db,userCollection,query,nil,&u)
	return &u,err
}

type UserRole struct {
	UserId			string
	RoleId 			string
	CreateUserId 	string
	CreateTime		int64  	//时间戳
	Role			int		//角色类型 （枚举表示）
	AppId			string 	//所属应用
	ModifyUserId 	string
}


type Application struct {
	Title			string	//应用（组件）名称
	AppId 			string
	Code 			int
	Desc			string  //项目描述
	CreateTime 		int64 	//创建时间
	ModifyTime 		int64
	ModifyUserId  	string
}
