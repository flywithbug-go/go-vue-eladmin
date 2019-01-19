package app_handler

import "vue-admin/web_server/server/handler/handler_common"

var Routers = []handler_common.GinHandleFunc{
	{
		Handler:    addApplicationHandler, //添加应用
		RouterType: handler_common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/app",
	},
	{
		Handler:    updateApplicationHandler,
		RouterType: handler_common.RouterTypeNeedAuth,
		Method:     "PUT",
		Route:      "/app",
	},
	{
		Handler:    getApplicationsHandler,
		RouterType: handler_common.RouterTypeNeedAuth,
		Method:     "GET",
		Route:      "/app/list",
	},
	{
		Handler:    getAllSimpleAppHandler,
		RouterType: handler_common.RouterTypeNeedAuth,
		Method:     "GET",
		Route:      "/app/list/simple",
	},
	{
		Handler:    removeApplicationHandler,
		RouterType: handler_common.RouterTypeNeedAuth,
		Method:     "DELETE",
		Route:      "/app",
	},
	{
		Handler:    addAppVersionHandler,
		RouterType: handler_common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/app/version/add",
	},
	{
		Handler:    removeAppVersionHandler,
		RouterType: handler_common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/app/version/remove",
	},
	{
		Handler:    getAppVersionListHandler,
		RouterType: handler_common.RouterTypeNeedAuth,
		Method:     "GET",
		Route:      "/app/version/list",
	},
	{
		Handler:    updateAppVersionHandler,
		RouterType: handler_common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/app/version/update",
	},
}
