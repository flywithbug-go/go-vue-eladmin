package handler

import (
	"doc-manager/web_server/common"
	"doc-manager/web_server/model"
	"net/http"
	"strings"

	"github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
)

type appPara struct {
	model.Application
	Time string `json:"time"`
}

func addApplicationHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	app := new(appPara)
	err := c.BindJSON(app)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid: "+err.Error())
		return
	}
	if app.BundleId == "" {
		aRes.SetErrorInfo(http.StatusBadRequest, "BundleId must fill")
		return
	}
	if app.Icon == "" {
		aRes.SetErrorInfo(http.StatusBadRequest, "Icon must fill")
		return
	}
	if app.Name == "" {
		aRes.SetErrorInfo(http.StatusBadRequest, "Name must fill")
		return
	}
	if len(app.Desc) < 10 {
		aRes.SetErrorInfo(http.StatusBadRequest, "Desc must fill")
		return
	}
	app.Owner = common.Account(c)
	err = app.ApplicationInsert()
	if err != nil {
		log4go.Error(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, err.Error())
		return
	}
	aRes.AddResponseInfo("app", app)
}

func getAllApplicationHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	userId := common.UserId(c)
	if strings.EqualFold(userId, "") {
		aRes.SetErrorInfo(http.StatusUnauthorized, "user not found")
		return
	}
	applist, err := model.FindALlApplications()
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusUnauthorized, "apps find error"+err.Error())
		return
	}
	aRes.AddResponseInfo("list", applist)
}
