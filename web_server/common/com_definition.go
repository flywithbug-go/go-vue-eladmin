package common

import (
	"github.com/gin-gonic/gin"
)

const (
	KeyUserToken = "Authorization"
	KeyJWTClaims = "_key_jwt_Claims"
	KeyUserAgent = "User-Agent"
	KeyUserId    = "user_id"
	KeyAccount   = "account"
)

func UserToken(ctx *gin.Context) string {
	token := ctx.GetHeader(KeyUserToken)
	return token
}

func UserId(ctx *gin.Context) string {
	o, ok := ctx.Get(KeyUserId)
	if !ok {
		return ""
	}
	userId, ok := o.(string)
	if !ok {
		return ""
	}
	return userId
}

func Account(ctx *gin.Context) string {
	o, ok := ctx.Get(KeyAccount)
	if !ok {
		return ""
	}
	userId, ok := o.(string)
	if !ok {
		return ""
	}
	return userId
}
