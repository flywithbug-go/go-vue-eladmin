package middleware

import (
	"net/http"
	"vue-admin/web_server/common"
	"vue-admin/web_server/core/jwt"
	"vue-admin/web_server/model"

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
		claims, err := jwt.ParseToken(token)
		if err != nil {
			delete(common.TokenValidMap, token)
			aRes.SetErrorInfo(http.StatusUnauthorized, err.Error())
			c.JSON(http.StatusUnauthorized, aRes)
			c.Abort()
			return
		}
		if !common.TokenValidMap[token] {
			_, err = model.FindLoginByToken(token)
			if err != nil {
				aRes.SetErrorInfo(http.StatusUnauthorized, "token无效，无权限访问")
				c.JSON(http.StatusUnauthorized, aRes)
				c.Abort()
				return
			}
			common.TokenValidMap[token] = true
		}
		c.Set(common.KeyContextUserId, claims.UserId)
		c.Set(common.KeyContextAccount, claims.Account)
	}
}
