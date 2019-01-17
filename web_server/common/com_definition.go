package common

import (
	"github.com/gin-gonic/gin"
)

const (
	KeyAuthorization    = "Authorization"
	KeyUserAgent        = "User-Agent"
	KeyContextPara      = "_key_ctx_para"
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

func Para(c *gin.Context) interface{} {
	para, ok := c.Get(KeyContextPara)
	if !ok {
		return nil
	}
	return para
}
