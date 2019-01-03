package role_handler

import (
	"net/http"
	"strconv"
	"vue-admin/web_server/model"

	"github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
)

func addPermissionHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(aRes.Code, aRes)
	}()
	p := model.Permission{}
	err := c.BindJSON(&p)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}

	if len(p.Code) == 0 {
		aRes.SetErrorInfo(http.StatusBadRequest, "code is null")
		return
	}
	if p.Type == 0 {
		aRes.SetErrorInfo(http.StatusBadRequest, "type should . 0")
		return
	}
	if len(p.Name) == 0 {
		aRes.SetErrorInfo(http.StatusBadRequest, "name not be null")
		return
	}
	if len(p.Description) == 0 {
		aRes.SetErrorInfo(http.StatusBadRequest, "desc should not be null")
		return
	}

	err = p.Insert()
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, "server invalid"+err.Error())
		return
	}
}

func getPermissionHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(aRes.Code, aRes)
	}()
	ids := c.Query("id")
	var mInfo = model.Permission{}
	id, _ := strconv.Atoi(ids)
	mInfo.Id = int64(id)
	mInfo, err := mInfo.FindOne()
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "invalid: "+err.Error())
		return
	}
	aRes.AddResponseInfo("permission", mInfo)
}

func updatePermissionHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(aRes.Code, aRes)
	}()
	para := model.Permission{}
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

func removePermissionHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(aRes.Code, aRes)
	}()
	//need id
	para := model.Permission{}
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
