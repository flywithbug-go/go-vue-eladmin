package handler

import (
	"doc-manager/core/jwt"
	log "github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
	"net/http"
)


func LoginHandler(ctx* gin.Context)  {
	aRes := NewResponse()
	defer func() {
		ctx.JSON(http.StatusOK, aRes)
	}()
	login := parameterModel{}
	err := ctx.BindJSON(&login)
	if err != nil {
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid" + err.Error())
		return
	}
	claims := jwt.NewCustomClaims("abc","abc")
	token ,err  := jwt.GenerateToken(claims)
	if err != nil {
		log.Error(err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, "token generate error" + err.Error())
		return
	}
	aRes.SetResponseDataInfo("token",token)
}
