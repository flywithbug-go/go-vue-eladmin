package role_handler

import "doc-manager/web_server/server/handler/common"

var FileRouters = []common.GinHandleFunc{
	{
		Handler:    addRoleHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "role/add",
	},
	{
		Handler:    addPermissionHandler,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "permission/add",
	},
}
