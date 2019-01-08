package user_handler

import "vue-admin/web_server/server/handler/common"

var UserRouters = []common.GinHandleFunc{

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
		Handler:    getUserInfoHandler, //获取当前用户信息
		RouterType: common.RouterTypeNeedAuth,
		Method:     "GET",
		Route:      "/user/info",
	},
	{
		Handler:    updateUserHandler, //更新当前用户信息
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/user/update",
	},
	{
		Handler:    getUserListInfoHandler, //获取所有用户
		RouterType: common.RouterTypeNeedAuth,
		Method:     "GET",
		Route:      "/user/list",
	},
}