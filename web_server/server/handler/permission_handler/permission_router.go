package permission_handler

import "vue-admin/web_server/server/handler/common"

var PermissionRouters = []common.GinHandleFunc{
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
