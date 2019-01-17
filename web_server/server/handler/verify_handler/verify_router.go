package verify_handler

import "vue-admin/web_server/server/handler/handler_common"

var Routers = []handler_common.GinHandleFunc{
	{
		Handler:    sendVerifyMailHanlder,
		RouterType: handler_common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/mail/verify",
	},
}
