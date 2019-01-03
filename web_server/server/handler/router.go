package handler

import (
	"doc-manager/web_server/server/handler/app_handler"
	"doc-manager/web_server/server/handler/common"
	"doc-manager/web_server/server/handler/file_handler"
	"doc-manager/web_server/server/handler/role_handler"

	"doc-manager/web_server/server/handler/user_handler"

	"doc-manager/web_server/server/middleware"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	routerList []common.GinHandleFunc
)

//host:port/auth_prefix/prefix/path
func RegisterRouters(r *gin.Engine, prefix string, authPrefix string) {
	r.NoRoute(NoRoute)
	jwtR := r.Group(prefix + authPrefix)
	jwtR.Use(middleware.JWTAuthMiddleware())

	addAllRouters()

	for _, v := range routerList {
		route := strings.ToLower(v.Route)
		switch v.RouterType {
		case common.RouterTypeNeedAuth:
			funcDoRouteNeedAuthRegister(strings.ToUpper(v.Method), route, v.Handler, jwtR)
		case common.RouterTypeNormal:
			route = strings.ToLower(prefix + route)
			funcDoRouteRegister(strings.ToUpper(v.Method), route, v.Handler, r)
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

func addAllRouters() {
	routerList = append(routerList, user_handler.UserRouters...)
	routerList = append(routerList, file_handler.FileRouters...)
	routerList = append(routerList, app_handler.AppRouters...)
	routerList = append(routerList, role_handler.FileRouters...)
}
