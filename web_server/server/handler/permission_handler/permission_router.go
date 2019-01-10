package permission_handler

import "vue-admin/web_server/server/handler/common"

var Routers = []common.GinHandleFunc{
	{
		Handler:    addPermissionHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/permission",
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
		Method:     "PUT",
		Route:      "/permission",
	},
	{
		Handler:    removePermissionHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "DELETE",
		Route:      "/permission",
	},
	{
		Handler:    getPermissionListHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "GET",
		Route:      "/permission/list",
	},
	{
		Handler:    getPermissionTreeHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "GET",
		Route:      "/permission/tree",
	},
}
