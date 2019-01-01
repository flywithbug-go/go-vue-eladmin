package middleware

import (
	"doc-manager/web_server/common"
	"doc-manager/web_server/core/jwt"
	"doc-manager/web_server/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

//JWTAuthMiddleware
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		aRes := model.NewResponse()
		//header拿token
		token := c.GetHeader(common.KeyAuthorization)
		if token == "" {
			//cookie拿token
			token, _ = c.Cookie(common.KeyAuthorization)
			if token == "" {
				aRes.SetErrorInfo(http.StatusUnauthorized, "请求未携带token，无权限访问")
				c.JSON(http.StatusUnauthorized, aRes)
				c.Abort()
				return
			}
		}
		l, err := model.FindLoginByToken(token)
		if err != nil {
			aRes.SetErrorInfo(http.StatusUnauthorized, "未查询到Token，无权限访问")
			c.JSON(http.StatusUnauthorized, aRes)
			c.Abort()
			return
		}
		if l.Status != 1 {
			aRes.SetErrorInfo(http.StatusUnauthorized, "授权已失效")
			c.JSON(http.StatusUnauthorized, aRes)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			model.UpdateLoginStatus(token, 0)
			if err == jwt.TokenExpired {
				aRes.SetErrorInfo(http.StatusUnauthorized, "授权已过期")
				c.JSON(http.StatusUnauthorized, aRes)
				c.Abort()
				return
			}
			aRes.SetErrorInfo(http.StatusUnauthorized, err.Error())
			c.JSON(http.StatusUnauthorized, aRes)
			c.Abort()
			return
		}
		user, err := model.FindByUserId(claims.UserId)
		if err != nil {
			aRes.SetErrorInfo(http.StatusUnauthorized, "未查询到User，无权限访问")
			c.JSON(http.StatusUnauthorized, aRes)
			c.Abort()
			return
		}
		c.Set(common.KeyContextUser, user)
		c.Set(common.KeyContextUserId, claims.UserId)
		c.Set(common.KeyContextAccount, claims.Account)
	}
}
