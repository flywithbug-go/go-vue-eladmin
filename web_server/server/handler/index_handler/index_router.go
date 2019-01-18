package index_handler

import "vue-admin/web_server/server/handler/handler_common"

var Routers = []handler_common.GinHandleFunc{
	{
		Handler:    indexHandler, //添加应用
		RouterType: handler_common.RouterTypeNormal,
		Method:     "GET",
		Route:      "/index",
	},
}
