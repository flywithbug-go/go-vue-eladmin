package file_handler

import "vue-admin/web_server/server/handler/handler_common"

var Routers = []handler_common.GinHandleFunc{
	{
		Handler:    uploadImageHandler, //上传图片
		RouterType: handler_common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/upload/image",
	},
	{
		Handler:    loadImageHandler, //加载图片
		RouterType: handler_common.RouterTypeNeedAuth,
		Method:     "GET",
		Route:      "/image/:path/:filename",
	},
}
