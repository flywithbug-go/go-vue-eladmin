package server

import (
	"doc-manager/server/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartWeb(port, staticPath string) {
	r := gin.New()
	r.Use(middleware.Logger(), gin.Recovery())
	r.Static("/", staticPath)
	err := r.Run(port)
	if err != nil {
		panic(fmt.Errorf("server启动失败 %s", err.Error()))
	}
}
