package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var(
	index = 0
)

// 系统状态信息
func IndexHandler(c *gin.Context) {
	aRes := NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	aRes.AddResponseInfo("moves","")
}
