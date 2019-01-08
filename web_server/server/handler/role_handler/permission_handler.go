package role_handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"vue-admin/web_server/common"
	"vue-admin/web_server/model"

	"github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
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
	if p.Type < model.PermissionTypeUndetermined || p.Type > model.PermissionTypeR {
		aRes.SetErrorInfo(http.StatusBadRequest, fmt.Sprintf("type should . %d", p.Type))
		return
	}
	if len(p.Name) == 0 {
		aRes.SetErrorInfo(http.StatusBadRequest, "name not be null")
		return
	}
	if len(p.Note) == 0 {
		aRes.SetErrorInfo(http.StatusBadRequest, "note should not be null")
		return
	}

	err = p.Insert()
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, "server invalid: "+err.Error())
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

func getPermissionListHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(aRes.Code, aRes)
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
	//if limit == 0 {
	//	limit = 10
	//}
	if page != 0 {
		page--
	}
	userId := common.UserId(c)
	if userId <= 0 {
		log4go.Info("user not found")
		aRes.SetErrorInfo(http.StatusUnauthorized, "user not found")
		return
	}
	query := bson.M{}
	var appV = model.Permission{}
	totalCount, _ := appV.TotalCount(query, nil)
	appList, err := appV.FindPageFilter(page, limit, query, nil, sort)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusUnauthorized, "app version list find error"+err.Error())
		return
	}
	aRes.AddResponseInfo("list", appList)
	aRes.AddResponseInfo("total", totalCount)
}
