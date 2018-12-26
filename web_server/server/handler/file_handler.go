package handler

import (
	"doc-manager/web_server/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

const  localImageFilePath  =  "./image/"

func imageUploadHandler(c *gin.Context)  {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK,aRes)
	}()






}