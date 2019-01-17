package verify_handler

import (
	"net/http"
	"vue-admin/web_server/common"
	"vue-admin/web_server/mail"
	"vue-admin/web_server/model"
	"vue-admin/web_server/model/model_verify"

	"github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
)

type mailVerifyPara struct {
	Mail string `json:"mail"`
}

func sendVerifyMailHanlder(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	verify := mailVerifyPara{}
	err := c.BindJSON(&verify)
	if err != nil {
		log4go.Info(common.XRequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "Param invalid"+err.Error())
		return
	}
	if !mail.MailVerify(verify.Mail) {
		aRes.SetErrorInfo(http.StatusBadRequest, "mail invalid")
		return
	}
	vCode, err := model_verify.GeneralVerifyData(verify.Mail)
	if err != nil {
		log4go.Info(common.XRequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, "invalid"+err.Error())
		return
	}
	err = mail.SendVerifyCode("", vCode, verify.Mail)
	if err != nil {
		log4go.Info(common.XRequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, err.Error())
		return
	}
	aRes.SetSuccess()

}
