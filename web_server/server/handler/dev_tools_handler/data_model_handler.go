package dev_tools_handler

import (
	"net/http"
	"vue-admin/web_server/common"
	"vue-admin/web_server/model"
	"vue-admin/web_server/model/model_dev_tools/model_data_model"
	"vue-admin/web_server/server/handler/check_permission"
	"vue-admin/web_server/server/handler/handler_common"

	"github.com/flywithbug/log4go"

	"github.com/gin-gonic/gin"
)

func addDataModelHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponseCode, aRes.Code)
		c.JSON(http.StatusOK, aRes)
	}()
	if check_permission.CheckNoPermission(c, model_data_model.DataModelPermissionCreate) {
		log4go.Info(handler_common.RequestId(c) + "has no permission")
		aRes.SetErrorInfo(http.StatusBadRequest, "has no permission")
		return
	}
	para := new(model_data_model.DataModel)
	err := c.BindJSON(para)
	if err != nil {
		log4go.Info(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	c.Set(common.KeyContextPara, para)
	if len(para.Name) == 0 {
		log4go.Info(handler_common.RequestId(c) + "name is nil")
		aRes.SetErrorInfo(http.StatusBadRequest, "name is nil")
		return
	}
	id, err := para.Insert()
	if err != nil {
		log4go.Info(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, "server invalid: "+err.Error())
		return
	}
	aRes.AddResponseInfo("id", id)
}

func updateDataModelHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponseCode, aRes.Code)
		c.JSON(http.StatusOK, aRes)
	}()
	if check_permission.CheckNoPermission(c, model_data_model.DataModelPermissionEdit) {
		log4go.Info(handler_common.RequestId(c) + "has no permission")
		aRes.SetErrorInfo(http.StatusBadRequest, "has no permission")
		return
	}
	para := new(model_data_model.DataModel)
	err := c.BindJSON(para)
	if err != nil {
		log4go.Info(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	c.Set(common.KeyContextPara, para)
	if para.Id == 0 {
		log4go.Info(handler_common.RequestId(c) + "id is 0")
		aRes.SetErrorInfo(http.StatusBadRequest, "id is 0")
		return
	}
	err = para.Update()
	if err != nil {
		log4go.Info(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "invalid: "+err.Error())
		return
	}
	aRes.SetSuccess()
}

func removeDataModelHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponseCode, aRes.Code)
		c.JSON(http.StatusOK, aRes)
	}()
	if check_permission.CheckNoPermission(c, model_data_model.DataModelPermissionDelete) {
		log4go.Info(handler_common.RequestId(c) + "has no permission")
		aRes.SetErrorInfo(http.StatusBadRequest, "has no permission")
		return
	}
	para := new(model_data_model.DataModel)
	err := c.BindJSON(para)
	if err != nil {
		log4go.Info(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	c.Set(common.KeyContextPara, para)
	if para.Id == 0 {
		log4go.Info(handler_common.RequestId(c) + "id is 0")
		aRes.SetErrorInfo(http.StatusBadRequest, "id is 0")
		return
	}
	err = para.Remove()
	if err != nil {
		log4go.Info(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "invalid: "+err.Error())
		return
	}
	aRes.SetSuccess()
}

func getDataModelHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponseCode, aRes.Code)
		c.JSON(http.StatusOK, aRes)
	}()
	if check_permission.CheckNoPermission(c, model_data_model.DataModelPermissionSelect) {
		log4go.Info(handler_common.RequestId(c) + "has no permission")
		aRes.SetErrorInfo(http.StatusBadRequest, "has no permission")
		return
	}
}
