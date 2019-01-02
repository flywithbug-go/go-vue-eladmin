package handler

import (
	"doc-manager/web_server/model"
	"net/http"

	"github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
)

func addRoleHandler(c *gin.Context) {

	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	user := model.User{}
	err := c.BindJSON(&user)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}

}
