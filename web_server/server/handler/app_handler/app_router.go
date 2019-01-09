package app_handler

import "vue-admin/web_server/server/handler/common"

var Routers = []common.GinHandleFunc{
	{
		Handler:    addApplicationHandler, //添加应用
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/app/add",
	},
	{
		Handler:    getApplicationsHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "GET",
		Route:      "/app/list",
	},
	{
		Handler:    getAllSimpleAppHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "GET",
		Route:      "/app/list/simple",
	},
	{
		Handler:    updateApplicationHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/app/update",
	},

	{
		Handler:    addAppVersionHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/app/version/add",
	},
	{
		Handler:    removeAppVersionHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/app/version/remove",
	},
	{
		Handler:    getAppVersionListHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "GET",
		Route:      "/app/version/list",
	},
	{
		Handler:    updateAppVersionHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/app/version/update",
	},
}
