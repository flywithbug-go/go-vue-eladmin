package handler

import (
	"doc-manager/common"
	"doc-manager/model"
	"github.com/flywithbug/log4go"
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
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	userId, _ := c.Get(common.KeyUserId)
	log4go.Info("index handler %s",userId)
	//log4go.Info("%v",model.AddAdminUser())
}
