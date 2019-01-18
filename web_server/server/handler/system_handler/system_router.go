package system_handler

import "vue-admin/web_server/server/handler/handler_common"

var Routers = []handler_common.GinHandleFunc{
	{
		Handler:    systemHandle,
		RouterType: handler_common.RouterTypeNeedAuth,
		Method:     "get",
		Route:      "/system",
	},
}
