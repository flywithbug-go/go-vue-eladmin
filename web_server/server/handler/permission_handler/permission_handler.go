package permission_handler

import (
	"net/http"
	"strconv"
	"strings"
	"vue-admin/web_server/common"
	"vue-admin/web_server/model"
	"vue-admin/web_server/model/model_permission"
	"vue-admin/web_server/server/handler/check_permission"
	"vue-admin/web_server/server/handler/handler_common"

	"gopkg.in/mgo.v2/bson"

	"github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
)

func addPermissionHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponseCode, aRes.Code)
		c.JSON(http.StatusOK, aRes)
	}()
	if check_permission.CheckNoPermission(c, model_permission.PPermissionCreate) {
		log4go.Info(handler_common.RequestId(c) + "has no permission")
		aRes.SetErrorInfo(http.StatusBadRequest, "has no permission")
		return
	}
	p := model_permission.Permission{}
	err := c.BindJSON(&p)
	if err != nil {
		log4go.Info(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	c.Set(common.KeyContextPara, p)

	id, err := p.Insert()
	if err != nil {
		log4go.Info(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, "server invalid: "+err.Error())
		return
	}
	aRes.AddResponseInfo("id", id)
}

func getPermissionHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponseCode, aRes.Code)
		c.JSON(http.StatusOK, aRes)
	}()
	if check_permission.CheckNoPermission(c, model_permission.PPermissionSelect) {
		log4go.Info(handler_common.RequestId(c) + "has no permission")
		aRes.SetErrorInfo(http.StatusBadRequest, "has no permission")
		return
	}
	ids := c.Query("id")
	id, _ := strconv.Atoi(ids)
	var mInfo = model_permission.Permission{}
	mInfo.Id = int64(id)
	mInfo, err := mInfo.FindOne(nil)
	if err != nil {
		log4go.Info(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "invalid: "+err.Error())
		return
	}
	aRes.AddResponseInfo("permission", mInfo)
}

func updatePermissionHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponseCode, aRes.Code)
		c.JSON(http.StatusOK, aRes)
	}()
	if check_permission.CheckNoPermission(c, model_permission.PPermissionEdit) {
		log4go.Info(handler_common.RequestId(c) + "has no permission")
		aRes.SetErrorInfo(http.StatusBadRequest, "has no permission")
		return
	}
	para := model_permission.Permission{}
	err := c.BindJSON(&para)
	if err != nil {
		log4go.Info(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	c.Set(common.KeyContextPara, para)

	err = para.Update()
	if err != nil {
		log4go.Info(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "invalid: "+err.Error())
		return
	}
	aRes.SetSuccess()
}

func removePermissionHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponseCode, aRes.Code)
		c.JSON(http.StatusOK, aRes)
	}()
	if check_permission.CheckNoPermission(c, model_permission.PPermissionDelete) {
		log4go.Info(handler_common.RequestId(c) + "has no permission")
		aRes.SetErrorInfo(http.StatusBadRequest, "has no permission")
		return
	}
	//need id
	para := model_permission.Permission{}
	err := c.BindJSON(&para)
	if err != nil {
		log4go.Info(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	c.Set(common.KeyContextPara, para)

	err = para.Remove()
	if err != nil {
		log4go.Info(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "invalid: "+err.Error())
		return
	}
	aRes.SetSuccess()

}

func getPermissionListHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponseCode, aRes.Code)
		c.JSON(http.StatusOK, aRes)
	}()
	if check_permission.CheckNoPermission(c, model_permission.PPermissionSelect) {
		log4go.Info(handler_common.RequestId(c) + "has no permission")
		aRes.SetErrorInfo(http.StatusBadRequest, "has no permission")
		return
	}
	var per = model_permission.Permission{}
	limit, _ := strconv.Atoi(c.Query("size"))
	page, _ := strconv.Atoi(c.Query("page"))
	sort := c.Query("sort")
	name := c.Query("name")
	if strings.EqualFold(sort, "-id") {
		sort = "-_id"
	} else if strings.EqualFold(sort, "+id") {
		sort = "+_id"
	} else if len(sort) == 0 {
		sort = "+_id"
	}
	if page != 0 {
		page--
	}
	query := bson.M{"pid": 0}
	if len(name) > 0 {
		query["name"] = bson.M{"$regex": name, "$options": "i"}
	}
	totalCount, _ := per.TotalCount(query, nil)
	list, err := per.FindPageFilter(page, limit, query, nil, sort)
	if err != nil {
		log4go.Info(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusUnauthorized, "apps find error"+err.Error())
		return
	}
	aRes.AddResponseInfo("list", list)
	aRes.AddResponseInfo("total", totalCount)

}

func getPermissionTreeHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponseCode, aRes.Code)
		c.JSON(http.StatusOK, aRes)
	}()
	if check_permission.CheckNoPermission(c, model_permission.PPermissionSelect) {
		log4go.Info(handler_common.RequestId(c) + "has no permission")
		aRes.SetErrorInfo(http.StatusBadRequest, "has no permission")
		return
	}
	var per = model_permission.Permission{}
	query := bson.M{"pid": 0}
	selector := bson.M{"_id": 1, "alias": 1}
	list, err := per.FindPageFilter(0, 0, query, selector)
	//list, err := per.FetchTreeList(selector)
	if err != nil {
		log4go.Info(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusUnauthorized, "app version list find error"+err.Error())
		return
	}
	aRes.AddResponseInfo("list", list)
}
