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
	totalCount, _ := model.TotalCountAppVersion(nil, nil)
	appList, err := model.FindPageAppVersionFilter(page, limit, nil, nil, sort)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusUnauthorized, "app version list find error"+err.Error())
		return
	}
	aRes.AddResponseInfo("list", appList)
	aRes.AddResponseInfo("total", totalCount)
}

func updateAppVersionHandler(c *gin.Context) {
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

}
