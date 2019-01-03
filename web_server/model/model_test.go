package model

import (
	"fmt"
	"testing"

	"vue-admin/web_server/core/mongo"
)

func TestUserFunctions(t *testing.T) {
	mongo.DialMgo("127.0.0.1:27017")

	user := new(User)
	user.RealName = "ori"
	user.Password = "pass"
	user.Account = "ori"
	user.Title = "CEO"
	user.Phone = "129"
	user.Email = "admin@admin.com"
	user.Sex = 1

	err := user.Insert()
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

}
