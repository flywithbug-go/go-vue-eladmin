package app_handler

import (
	"net/http"
	"strconv"
	"strings"
	"vue-admin/web_server/common"
	"vue-admin/web_server/model"
	"vue-admin/web_server/model/model_app"
	"vue-admin/web_server/server/handler/check_permission"

	"gopkg.in/mgo.v2/bson"

	"github.com/flywithbug/log4go"

	"github.com/gin-gonic/gin"
)

func addAppVersionHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	if check_permission.CheckNoPermission(c, model_app.APPlicationPermissionCreate) {
		log4go.Info(common.XRequestId(c) + "has no permission")
		aRes.SetErrorInfo(http.StatusOK, "has no permission")
		return
	}

	appV := new(model_app.AppVersion)
	err := c.BindJSON(appV)
	if err != nil {
		log4go.Info(common.XRequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid: "+err.Error())
		return
	}
	err = appV.Insert()
	if err != nil {
		log4go.Info(common.XRequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, "para invalid: "+err.Error())
		return
	}
	aRes.AddResponseInfo("app", appV)
}

func updateAppVersionHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	if check_permission.CheckNoPermission(c, model_app.APPlicationPermissionEdit) {
		log4go.Info(common.XRequestId(c) + "has no permission")
		aRes.SetErrorInfo(http.StatusOK, "has no permission")
		return
	}

	appV := new(model_app.AppVersion)
	err := c.BindJSON(appV)
	if err != nil {
		log4go.Info(common.XRequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid: "+err.Error())
		return
	}
	err = appV.Update()
	if err != nil {
		log4go.Info(common.XRequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, "para invalid: "+err.Error())
		return
	}
	aRes.SetSuccess()
}

func getAppVersionListHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	if check_permission.CheckNoPermission(c, model_app.APPlicationPermissionSelect) {
		log4go.Info(common.XRequestId(c) + "has no permission")
		aRes.SetErrorInfo(http.StatusOK, "has no permission")
		return
	}

	appId, _ := strconv.Atoi(c.Query("app_id"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	sort := c.Query("sort")
	if strings.EqualFold(sort, "-id") {
		sort = "-_id"
	} else if strings.EqualFold(sort, "+id") {
		sort = "+_id"
	} else if len(sort) == 0 {
		sort = "+_id"
	}
	//if limit == 0 {
	//	limit = 10
	//}
	if page != 0 {
		page--
	}
	userId := common.UserId(c)
	if userId <= 0 {
		log4go.Info(common.XRequestId(c) + "user not found")
		aRes.SetErrorInfo(http.StatusUnauthorized, "user not found")
		return
	}
	query := bson.M{}
	if appId > 0 {
		query = bson.M{"app_id": appId}
	}
	var appV = model_app.AppVersion{}
	totalCount, _ := appV.TotalCount(query, nil)
	appList, err := appV.FindPageFilter(page, limit, query, nil, sort)
	if err != nil {
		log4go.Info(common.XRequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusUnauthorized, "app version list find error"+err.Error())
		return
	}
	aRes.AddResponseInfo("list", appList)
	aRes.AddResponseInfo("total", totalCount)
}

func removeAppVersionHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	if check_permission.CheckNoPermission(c, model_app.APPlicationPermissionDelete) {
		log4go.Info(common.XRequestId(c) + "has no permission")
		aRes.SetErrorInfo(http.StatusOK, "has no permission")
		return
	}
	appV := model_app.AppVersion{}
	err := c.BindJSON(&appV)
	if err != nil {
		log4go.Info(common.XRequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid: "+err.Error())
		return
	}
	err = appV.Remove()
	if err != nil {
		log4go.Info(common.XRequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, "para invalid: "+err.Error())
		return
	}
	aRes.SetSuccess()
}
