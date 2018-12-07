package server

import (
	"doc-manager/logger"
	"doc-manager/server/handler"
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartApi(port ,staticPath string, rPrefix []string) {
	r := gin.New()
	r.Use(logger.Logger(),gin.Recovery())
	r.Static("/index.html",staticPath)
	handler.RegisterRouters(r,rPrefix)
	err := r.Run(port)
	if err != nil {
		panic(fmt.Errorf("server启动失败 %s", err.Error()))
	}
}
