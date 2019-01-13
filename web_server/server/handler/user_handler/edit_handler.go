package user_handler

import (
	"net/http"
	"strings"
	"vue-admin/web_server/common"
	"vue-admin/web_server/model"
	"vue-admin/web_server/model/model_user"
	"vue-admin/web_server/model/model_verify"

	"github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
)

type ParaUser struct {
	Password    string `json:"password"`
	OldPassword string `json:"old_password"`
	Code        string `json:"code"`
	Mail        string `json:"mail"`
}

func validPasswordHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	password := c.Query("password")
	if len(password) == 0 {
		log4go.Info("password need")
		aRes.SetErrorInfo(http.StatusBadRequest, "password need")
		return
	}
	user := new(model_user.User)
	user.Password = password
	user.Id = common.UserId(c)
	if !user.CheckPassword() {
		log4go.Info("password not right")
		aRes.SetErrorInfo(http.StatusBadRequest, "password not right")
		return
	}
	aRes.SetSuccess()
}

func updatePasswordHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	para := new(ParaUser)
	err := c.BindJSON(para)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	if strings.EqualFold(para.Password, para.OldPassword) {
		aRes.SetErrorInfo(http.StatusBadRequest, "password not changed")
		return
	}
	if len(para.Password) == 0 || len(para.OldPassword) == 0 {
		log4go.Info("para not right")
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid")
		return
	}
	user := new(model_user.User)
	user.Password = para.OldPassword
	user.Id = common.UserId(c)
	if !user.CheckPassword() {
		log4go.Info("password not right")
		aRes.SetErrorInfo(http.StatusBadRequest, "password not right")
		return
	}
	user.Password = para.Password
	err = user.UpdatePassword()
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, "system error: "+err.Error())
		return
	}
	aRes.SetSuccess()
}

func updateMailHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	para := new(ParaUser)
	err := c.BindJSON(para)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	if len(para.Mail) == 0 || len(para.Password) == 0 || len(para.Code) == 0 {
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid")
		return
	}

	if !model_verify.CheckVerify(para.Mail, para.Code) {
		aRes.SetErrorInfo(http.StatusBadRequest, "code not right")
		return
	}
	var user model_user.User
	user.Id = common.UserId(c)
	user.Password = para.Password
	if !user.CheckPassword() {
		aRes.SetErrorInfo(http.StatusBadRequest, "password not right")
		return
	}
	user.Email = para.Mail
	user.UpdateMail()
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, "system error: "+err.Error())
		return
	}
	aRes.SetSuccess()
}
