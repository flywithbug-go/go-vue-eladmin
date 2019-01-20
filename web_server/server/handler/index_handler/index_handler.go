package index_handler

import (
	"net/http"
	"time"
	"vue-admin/web_server/common"
	"vue-admin/web_server/model"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponseCode, aRes.Code)
		c.JSON(http.StatusOK, aRes)
	}()
	aRes.SetSuccess()
	time.Sleep(time.Second * 10)
}
