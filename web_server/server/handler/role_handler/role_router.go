package role_handler

import "doc-manager/web_server/server/handler/common"

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
		Handler:    deleteRoleHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/role/delete",
	},
	{
		Handler:    addPermissionHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/permission/add",
	},

	{
		Handler:    updatePermissionHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/permission/update",
	},

	{
		Handler:    deletePermissionHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/permission/delete",
	},
}
