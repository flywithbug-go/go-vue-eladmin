package role_handler

import (
	"doc-manager/web_server/model"
	"net/http"
	"strconv"

	"github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
)

type paraRole struct {
	model.Role
}

func addRoleHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	para := paraRole{}
	err := c.BindJSON(&para)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	err = para.Insert()
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, "server invalid"+err.Error())
		return
	}
}

func getRoleHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	ids := c.Query("id")
	var role = model.Role{}
	id, _ := strconv.Atoi(ids)
	role.Id = int64(id)
	role, err := role.FindOne()
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "invalid: "+err.Error())
		return
	}
	aRes.AddResponseInfo("role", role)
}

func updateRoleHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	para := model.Role{}
	err := c.BindJSON(&para)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	err = para.Update()
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "invalid: "+err.Error())
		return
	}
}

func deleteRoleHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	//need id
	para := model.Role{}
	err := c.BindJSON(&para)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	err = para.Remove()
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "invalid: "+err.Error())
		return
	}
}
