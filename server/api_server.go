package server

import (
	"doc-manager/server/handler"
	"doc-manager/server/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartApi(port string, rPrefix string,auth_prefix string) {
	r := gin.New()
	r.Use(middleware.Logger(),gin.Recovery())
	r.Use(middleware.CookieMiddleware())
	handler.RegisterRouters(r,rPrefix,auth_prefix)
	err := r.Run(port)
	if err != nil {
		panic(fmt.Errorf("server启动失败 %s", err.Error()))
	}
}
