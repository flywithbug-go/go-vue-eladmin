package model

import (
	"fmt"
	"testing"

	"vue-admin/web_server/core/mongo"
)

func TestUserFunctions(t *testing.T) {
	mongo.DialMgo("doc:doc11121014a@118.89.108.25:27017/docmanager")

	user := new(User)
	user.Password = "admin"
	user.Account = "admin"
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
