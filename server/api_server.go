package server

import (
	"doc-manager/logger"
	"doc-manager/server/handler"
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartApi(port string, rPrefix string,auth_prefix string) {
	r := gin.New()
	r.Use(logger.Logger(),gin.Recovery())
	handler.RegisterRouters(r,rPrefix,auth_prefix)
	err := r.Run(port)
	if err != nil {
		panic(fmt.Errorf("server启动失败 %s", err.Error()))
	}
}
