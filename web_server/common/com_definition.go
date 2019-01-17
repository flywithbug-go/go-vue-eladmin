package common

import (
	"github.com/gin-gonic/gin"
)

const (
	KeyAuthorization    = "Authorization"
	KeyUserAgent        = "User-Agent"
	KeyContextUserId    = "_key_ctx_userId_"
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
