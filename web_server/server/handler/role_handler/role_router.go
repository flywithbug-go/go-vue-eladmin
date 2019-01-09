package role_handler

import "vue-admin/web_server/server/handler/common"

var Routers = []common.GinHandleFunc{
	{
		Handler:    addRoleHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/role/add",
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
		Method:     "POST",
		Route:      "/role/update",
	},
	{
		Handler:    removeRoleHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/role/remove",
	},
}
