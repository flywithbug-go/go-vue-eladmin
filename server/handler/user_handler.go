package handler

import (
	"doc-manager/common"
	"doc-manager/core/jwt"
	"doc-manager/model"
	"net/http"
	"strings"

	log "github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	user := new(model.User)
	err := c.BindJSON(user)
	if err != nil {
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	err = user.CheckLogin(user.Account, user.Password)
	if err != nil {
		log.Error(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "account or password not right")
		return
	}
	claims := jwt.NewCustomClaims(user.UserId, user.Account)
	token, err := jwt.GenerateToken(claims)
	if err != nil {
		log.Error(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "token generate error"+err.Error())
		return
	}
	userAg := c.GetHeader(common.KeyUserAgent)
	_, err = model.UserLogin(user.UserId, userAg, token, c.ClientIP())
	if err != nil {
		log.Error(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "token generate error"+err.Error())
		return
	}
	aRes.SetResponseDataInfo("token", token)
}

func RegisterHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(aRes.Code, aRes)
	}()
	user := new(model.User)
	err := c.BindJSON(user)
	if err != nil {
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	if user.Account == "" {
		aRes.SetErrorInfo(http.StatusBadRequest, "account can not be nil")
		return
	}
	if user.Password == "" {
		aRes.SetErrorInfo(http.StatusBadRequest, "Password can not be nil")
		return
	}
	if user.Email == "" {
		aRes.SetErrorInfo(http.StatusBadRequest, "email can not be nil")
		return
	}
	err = user.Insert()
	if err != nil {
		aRes.SetErrorInfo(http.StatusBadRequest, err.Error())
		return
	}
	aRes.AddResponseInfo("user", user)
}

func LogoutHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(aRes.Code, aRes)
	}()
	token := common.UserToken(c)
	if token == "" {
		aRes.SetErrorInfo(http.StatusBadRequest, "token not found")
		return
	}
	err := model.UpdateLoginStatus(token, model.StatusLogout)
	if err != nil {
		aRes.SetErrorInfo(http.StatusBadRequest, err.Error())
		return
	}
	aRes.SetSuccessInfo(http.StatusOK, "success")
}

func GetUserInfoHandler(c *gin.Context)  {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(aRes.Code, aRes)
	}()
	userId := common.UserId(c)
	if strings.EqualFold(userId,"") {
		aRes.SetErrorInfo(http.StatusUnauthorized, "user not found")
		return
	}
	user ,err := model.FindByUserId(userId)
	if err != nil {
		aRes.SetErrorInfo(http.StatusUnauthorized, "user not found:" +err.Error())
		return
	}
	aRes.AddResponseInfo("user",user)
}
