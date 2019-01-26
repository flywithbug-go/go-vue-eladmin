package index_handler

import (
	"net/http"
	"vue-admin/web_server/common"
	"vue-admin/web_server/model"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponse, aRes)
		c.JSON(http.StatusOK, aRes)
	}()
	aRes.SetSuccess()

}
