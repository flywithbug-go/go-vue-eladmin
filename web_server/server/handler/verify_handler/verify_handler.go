package verify_handler

import (
	"net/http"
	"strconv"
	"vue-admin/web_server/common"
	"vue-admin/web_server/email"
	"vue-admin/web_server/model"
	"vue-admin/web_server/model/model_verify"
	"vue-admin/web_server/server/handler/handler_common"

	"gopkg.in/mgo.v2/bson"

	"github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
)

type mailVerifyPara struct {
	Mail string `json:"mail"`
}

func sendVerifyMailHanlder(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponseCode, aRes.Code)
		c.JSON(http.StatusOK, aRes)
	}()
	verify := mailVerifyPara{}
	err := c.BindJSON(&verify)
	if err != nil {
		log4go.Info(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "Param invalid"+err.Error())
		return
	}
	c.Set(common.KeyContextPara, verify)

	if !email.MailVerify(verify.Mail) {
		aRes.SetErrorInfo(http.StatusBadRequest, "mail invalid")
		return
	}
	vCode, err := model_verify.GeneralVerifyData(verify.Mail)
	if err != nil {
		log4go.Info(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, "invalid"+err.Error())
		return
	}
	err = email.SendVerifyCode("", vCode, verify.Mail)
	if err != nil {
		log4go.Info(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, err.Error())
		return
	}
	aRes.SetSuccess()
}

func getVerifyListHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponseCode, aRes.Code)
		c.JSON(http.StatusOK, aRes)
	}()
	size, _ := strconv.Atoi(c.Query("size"))
	page, _ := strconv.Atoi(c.Query("page"))
	value := c.Query("value")
	code := c.Query("code")
	scenes := c.Query("scenes")
	if size == 0 {
		size = 10
	}
	if page != 0 {
		page--
	}
	query := bson.M{}
	if len(value) > 0 {
		query["value"] = value
	}

	if len(code) > 0 {
		query["code"] = code
	}

	if len(scenes) > 0 {
		query["scenes"] = scenes
	}
	var v = model_verify.VerificationCode{}
	totalCount, _ := v.TotalCount(query, nil)
	results, err := v.FindPageFilter(page, size, query, nil, "-_id")
	if err != nil {
		log4go.Error(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, "list find error"+err.Error())
		return
	}
	aRes.AddResponseInfo("list", results)
	aRes.AddResponseInfo("total", totalCount)
}
