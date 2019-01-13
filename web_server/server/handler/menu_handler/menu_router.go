package menu_handler

import "vue-admin/web_server/server/handler/common"

var Routers = []common.GinHandleFunc{
	{
		Handler:    addMenuHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/menu",
	},
	{
		Handler:    getMenuHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "GET",
		Route:      "/menu",
	},
	{
		Handler:    updateMenuHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "PUT",
		Route:      "/menu",
	},
	{
		Handler:    removeMenuHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "DELETE",
		Route:      "/menu",
	},
	{
		Handler:    getRoleListHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "GET",
		Route:      "/menu/list",
	},
	{
		Handler:    getRoleTreeHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "GET",
		Route:      "/menu/tree",
	},
}
