package handler

import (
	"strings"

	"doc-manager/web_server/server/middleware"

	"github.com/gin-gonic/gin"
)

type stateType int

const (
	routerTypeNormal stateType = iota
	routerTypeNeedAuth
)

type ginHandleFunc struct {
	handler    gin.HandlerFunc
	routerType stateType
	method     string
	route      string
}

//host:port/auth_prefix/prefix/path
func RegisterRouters(r *gin.Engine, prefix string, authPrefix string) {
	jwtR := r.Group(prefix + authPrefix)
	jwtR.Use(middleware.JWTAuthMiddleware())
	for _, v := range routers {
		route := strings.ToLower(v.route)
		switch v.routerType {
		case routerTypeNeedAuth:
			funcDoRouteNeedAuthRegister(strings.ToUpper(v.method), route, v.handler, jwtR)
		case routerTypeNormal:
			route = strings.ToLower(prefix + route)
			funcDoRouteRegister(strings.ToUpper(v.method), route, v.handler, r)
		}
	}
}

//使用jwt过滤的routerType==routerTypeNeedAuth
func funcDoRouteNeedAuthRegister(method, route string, handler gin.HandlerFunc, jwtR *gin.RouterGroup) {
	switch method {
	case "POST":
		jwtR.POST(route, handler)
	case "PUT":
		jwtR.PUT(route, handler)
	case "HEAD":
		jwtR.HEAD(route, handler)
	case "DELETE":
		jwtR.DELETE(route, handler)
	case "OPTIONS":
		jwtR.OPTIONS(route, handler)
	default:
		jwtR.GET(route, handler)
	}
}

//普通routerType==routerTypeNormal
func funcDoRouteRegister(method, route string, handler gin.HandlerFunc, r *gin.Engine) {
	switch method {
	case "POST":
		r.POST(route, handler)
	case "PUT":
		r.PUT(route, handler)
	case "HEAD":
		r.HEAD(route, handler)
	case "DELETE":
		r.DELETE(route, handler)
	case "OPTIONS":
		r.OPTIONS(route, handler)
	default:
		r.GET(route, handler)
	}
}

var routers = []ginHandleFunc{
	{
		handler:    registerHandler,
		routerType: routerTypeNormal,
		method:     "POST",
		route:      "/register",
	},
	{
		handler:    loginHandler,
		routerType: routerTypeNormal,
		method:     "POST",
		route:      "/login",
	},
	{
		handler:    logoutHandler,
		routerType: routerTypeNeedAuth,
		route:      "/logout",
		method:     "POST",
	},
	{
		handler:    getUserInfoHandler, //获取当前用户信息
		routerType: routerTypeNeedAuth,
		method:     "GET",
		route:      "/user/info",
	},
	{
		handler:    updateUserHandler, //更新当前用户信息
		routerType: routerTypeNeedAuth,
		method:     "POST",
		route:      "/user/update",
	},
	{
		handler:    getUserListInfoHandler, //获取所有用户
		routerType: routerTypeNeedAuth,
		method:     "GET",
		route:      "/user/list",
	},
	{
		handler:    uploadImageHandler, //上传图片
		routerType: routerTypeNeedAuth,
		method:     "POST",
		route:      "/upload/image",
	},
	{
		handler:    loadImageHandler, //加载图片
		routerType: routerTypeNormal,
		method:     "GET",
		route:      "/image/:path/:filename",
	},
	{
		handler:    addApplicationHandler, //添加应用
		routerType: routerTypeNeedAuth,
		method:     "POST",
		route:      "/app/add",
	},
	{
		handler:    getApplicationsHandler,
		routerType: routerTypeNeedAuth,
		method:     "GET",
		route:      "/app/list",
	},
	{
		handler:    updateApplicationHandler,
		routerType: routerTypeNeedAuth,
		method:     "POST",
		route:      "/app/update",
	},
	{
		handler:    addAppVersionHandler,
		routerType: routerTypeNeedAuth,
		method:     "POST",
		route:      "/app/version/add",
	},
	{
		handler:    getAppVersionListHandler,
		routerType: routerTypeNeedAuth,
		method:     "GET",
		route:      "/app/version/list",
	},
	{
		handler:    updateAppVersionHandler,
		routerType: routerTypeNeedAuth,
		method:     "POST",
		route:      "/app/version/update",
	},
	{
		handler:    getAllSimpleAppHandler,
		routerType: routerTypeNeedAuth,
		method:     "GET",
		route:      "app/list/simple",
	},
}
