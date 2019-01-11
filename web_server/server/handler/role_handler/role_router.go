package role_handler

import "vue-admin/web_server/server/handler/common"

var Routers = []common.GinHandleFunc{
	{
		Handler:    addRoleHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/role",
	},
	{
		Handler:    getRoleHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "GET",
		Route:      "/role",
	},
	{
		Handler:    updateRoleHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "PUT",
		Route:      "/role",
	},
	{
		Handler:    removeRoleHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "DELETE",
		Route:      "/role",
	},
	{
		Handler:    getRoleListHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "GET",
		Route:      "/role/list",
	},
}
