package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartWeb(port ,staticPath string) {
	r := gin.Default()
	//r.Use(logger.Logger(),gin.Recovery())
	r.Static("/",staticPath)
	err := r.Run(port)
	if err != nil {
		panic(fmt.Errorf("server启动失败 %s", err.Error()))
	}
}
