package handler



import (
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
)


var routers = map[string]gin.HandlerFunc{
	"GET   		/": 							IndexHandler, //系统状态
	"GET   		/index": 							IndexHandler, //系统状态

	//"GET       /version":                      versionHandler, //获取应用最新版本
}

func RegisterRouters(r *gin.Engine, prefixs []string){
	dup := make(map[string]bool)
	for _, p := range prefixs {
		dup[p] = true
	}
	if len(dup) == 0 {
		dup[""] = true
	}
	for router, handler := range routers {
		method ,path := regexpRouters(router)
		for  k := range dup {
			funcDoRouteRegister(method,strings.ToLower(k+path),handler,r)//path 全小写
		}
	}
}

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

var routerRe = regexp.MustCompile(`([A-Z]+)[^/]*(/[a-zA-Z/:]+)`)
func regexpRouters(router string) (method,path string) {
	match := routerRe.FindAllStringSubmatch(router, -1)
	if len(match)<1 {
		return
	}
	return match[0][1],match[0][2]
}
