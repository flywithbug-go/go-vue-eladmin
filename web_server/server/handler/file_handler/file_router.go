package file_handler

import "vue-admin/web_server/server/handler/common"

var Routers = []common.GinHandleFunc{
	{
		Handler:    uploadImageHandler, //上传图片
		RouterType: common.RouterTypeNeedAuth,
		Method:     "POST",
		Route:      "/upload/image",
	},
	{
		Handler:    loadImageHandler, //加载图片
		RouterType: common.RouterTypeNormal,
		Method:     "GET",
		Route:      "/image/:path/:filename",
	},
}
