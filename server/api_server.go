package server

import (
	"doc-manager/server/handler"
	"doc-manager/server/middleware"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartApi(port string, rPrefix string, authPrefix string) {
	r := gin.New()
	r.Use(middleware.Logger(), gin.Recovery())
	cors.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Content-Length"},
		ExposeHeaders:    []string{},
		MaxAge:           12 * time.Hour,
		AllowCredentials: false,
	}))
	r.Use(middleware.CookieMiddleware())
	handler.RegisterRouters(r, rPrefix, authPrefix)
	err := r.Run(port)
	if err != nil {
		panic(fmt.Errorf("server启动失败 %s", err.Error()))
	}
}
