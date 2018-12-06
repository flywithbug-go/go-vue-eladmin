package server

import (
	"doc_manager/logger"
	"doc_manager/server/handler"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Start(port string, rPrefix []string) {
	r := gin.New()
	r.Use(logger.Logger())
	handler.RegisterRouters(r,nil)
	err := r.Run(port)
	if err != nil {
		panic(fmt.Errorf("server启动失败 %s", err.Error()))
	}
}
