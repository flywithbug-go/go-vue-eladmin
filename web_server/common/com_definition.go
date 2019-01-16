package common

import (
	"github.com/gin-gonic/gin"
)

const (
	KeyAuthorization    = "Authorization"
	KeyUserAgent        = "User-Agent"
	KeyContextUserId    = "_key_ctx_userId_"
	KeyContextUsername  = "_key_ctx_username_"
	KeyContextRequestId = "X-Reqid"
)

func UserToken(ctx *gin.Context) string {
	token := ctx.GetHeader(KeyAuthorization)
	return token
}

func UserId(ctx *gin.Context) int64 {
	o, ok := ctx.Get(KeyContextUserId)
	if !ok {
		return -1
	}
	userId, ok := o.(int64)
	if !ok {
		return -1
	}
	return userId
}

func Username(ctx *gin.Context) string {
	o, ok := ctx.Get(KeyContextUsername)
	if !ok {
		return ""
	}
	userId, ok := o.(string)
	if !ok {
		return ""
	}
	return userId
}

func XRequestId(ctx *gin.Context) string {
	o, ok := ctx.Get(KeyContextRequestId)
	if !ok {
		return ""
	}
	requestId, ok := o.(string)
	if !ok {
		return ""
	}
	return requestId
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
//}
