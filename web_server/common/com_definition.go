package common

import (
	"github.com/gin-gonic/gin"
)

const (
	KeyUserToken       = "Authorization"
	KeyJWTClaims       = "_key_jwt_Claims"
	KeyUserAgent       = "User-Agent"
	KeyContextUserId   = "_key_ctx_userId_"
	KeyContextAccount  = "_key_ctx_account_"
	KeyContextUser     = "_key_ctx_user_"
	KeyContextUserDBID = "_key_ctx_user_DBID"
)

func UserToken(ctx *gin.Context) string {
	token := ctx.GetHeader(KeyUserToken)
	return token
}

func UserId(ctx *gin.Context) string {
	o, ok := ctx.Get(KeyContextUserId)
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
	o, ok := ctx.Get(KeyContextAccount)
	if !ok {
		return ""
	}
	userId, ok := o.(string)
	if !ok {
		return ""
	}
	return userId
}

//func User(ctx *gin.Context) (*model.User, bool) {
//	o, ok := ctx.Get(KeyContextUser)
//	if !ok {
//		return nil, false
//	}
//	aUser, ok := o.(*model.User)
//	if !ok {
//		return nil, false
//	}
//	return aUser, true
//
//}
