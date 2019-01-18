package server

import (
	"fmt"
	"net/http"
	"time"
	"vue-admin/web_server/server/handler"
	"vue-admin/web_server/server/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func StartServer(port, staticPath, rPrefix, authPrefix string) {
	r := gin.New()
	r.Use(middleware.Logger(), gin.Recovery())
	r.Use(static.Serve("/", static.LocalFile(staticPath, true)))
	r.LoadHTMLGlob(staticPath + "/index.html")
	r.NoRoute(NoRoute)
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
	//err := r.Run(port)
	s := &http.Server{
		Addr:           port,
		Handler:        r,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(fmt.Errorf("server启动失败 %s", err.Error()))
	}
}
