package user_handler

import "vue-admin/web_server/server/handler/common"

var Routers = []common.GinHandleFunc{
	{
		Handler:    registerHandler,
		RouterType: common.RouterTypeNormal,
		Method:     "POST",
		Route:      "/register",
	},
	{
		Handler:    loginHandler,
		RouterType: common.RouterTypeNormal,
		Method:     "POST",
		Route:      "/login",
	},
	{
		Handler:    logoutHandler,
		RouterType: common.RouterTypeNeedAuth,
		Route:      "/logout",
		Method:     "POST",
	},
	{
		Handler:    getUserInfoHandler, //获取userId对应的用户信息
		RouterType: common.RouterTypeNeedAuth,
		Method:     "GET",
		Route:      "/user/info",
	},
	{
		Handler:    updateUserHandler, //更新当前用户信息
		RouterType: common.RouterTypeNeedAuth,
		Method:     "PUT",
		Route:      "/user",
	},
	{
		Handler:    addUserHandler, //添加用户当前用户信息
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/user",
	},
	{
		Handler:    deleteUserHandler, //删除当前用户信息
		RouterType: common.RouterTypeNeedAuth,
		Method:     "DELETE",
		Route:      "/user",
	},
	{
		Handler:    getUserListInfoHandler, //获取所有用户
		RouterType: common.RouterTypeNeedAuth,
		Method:     "GET",
		Route:      "/user/list",
	},
}
