package handler

import (
	"doc-manager/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

var(
	index = 0
)
var user model.User
var login model.Login
// 系统状态信息
func IndexHandler(c *gin.Context) {
	aRes := NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	//users, _ := user.FindAll()
	//aRes.AddResponseInfo("users",users)
	l,_ := model.UserLogin("abc","aaa")
	aRes.AddResponseInfo("token",l.Token)
}
