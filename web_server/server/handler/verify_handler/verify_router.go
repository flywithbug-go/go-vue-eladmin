package verify_handler

import "vue-admin/web_server/server/handler/common"

var Routers = []common.GinHandleFunc{
	{
		Handler:    sendVerifyMailHanlder,
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/mail",
	},
}
