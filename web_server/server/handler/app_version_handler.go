package handler

import (
	"doc-manager/web_server/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addAppVersionHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()

}
