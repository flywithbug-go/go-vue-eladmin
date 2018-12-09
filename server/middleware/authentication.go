package middleware

import (
	"doc-manager/common"
	"doc-manager/core/jwt"
	"github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTAuthMiddleware() gin.HandlerFunc  {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(common.KeyUserToken)
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}
		log4go.Info("get token: %s",token)
		j := jwt.NewJWT()
		claims ,err := j.ParseToken(token)
		if err != nil {
			if err == jwt.TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		c.Set(common.KeyJWTClaims,claims)
	}
}
