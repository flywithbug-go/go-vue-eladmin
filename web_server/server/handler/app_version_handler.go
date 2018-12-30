package handler

import (
	"doc-manager/web_server/model"
	"net/http"

	"github.com/flywithbug/log4go"

	"github.com/gin-gonic/gin"
)

type appVersionPara struct {
	model.AppVersion
}

func addAppVersionHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	appV := new(appVersionPara)
	err := c.BindJSON(appV)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid: "+err.Error())
		return
	}
	err = appV.Insert()
	if err != nil {
		aRes.SetErrorInfo(http.StatusInternalServerError, "db invalid: "+err.Error())
		return
	}
	aRes.SetSuccess()
}
