package handler

import (
	"doc-manager/web_server/common"
	"doc-manager/web_server/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
)

type appPara struct {
	model.Application
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

	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	sort := c.Query("sort")
	if strings.EqualFold(sort, "+id") {
		sort = "+_id"
	} else if strings.EqualFold(sort, "-id") {
		sort = "-_id"
	}
	if limit == 0 {
		limit = 10
	}
	if page != 0 {
		page--
	}
	userId := common.UserId(c)
	if strings.EqualFold(userId, "") {
		aRes.SetErrorInfo(http.StatusUnauthorized, "user not found")
		return
	}
	totalCount, _ := model.TotalCountApplication()
	applist, err := model.FindPageApplications(page, limit, sort)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusUnauthorized, "apps find error"+err.Error())
		return
	}
	aRes.AddResponseInfo("list", applist)
	aRes.AddResponseInfo("total", totalCount)
}

func updateApplicationHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	app := new(model.Application)
	err := c.BindJSON(app)
	if err != nil {
		log4go.Error(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid: "+err.Error())
		return
	}
	if app.AppId == "" {
		aRes.SetErrorInfo(http.StatusBadRequest, "appId invalid: ")
		return
	}

	err = model.UpdateApplication(app)
	if err != nil {
		log4go.Error(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "update failed: "+err.Error())
		return
	}
	aRes.SetSuccessInfo(http.StatusOK, "success")
}
