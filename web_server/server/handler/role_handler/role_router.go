package role_handler

import "vue-admin/web_server/server/handler/common"

var FileRouters = []common.GinHandleFunc{
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
	{
		Handler:    addPermissionHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/permission/add",
	},
	{
		Handler:    getPermissionHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "GET",
		Route:      "/permission",
	},
	{
		Handler:    updatePermissionHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/permission/update",
	},
	{
		Handler:    removePermissionHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/permission/remove",
	},
	{
		Handler:    getPermissionListHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "GET",
		Route:      "/permission/list",
	},
}
