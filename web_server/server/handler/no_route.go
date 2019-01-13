package handler

import (
	"net/http"
	"strings"
	"vue-admin/web_server/model"

	"github.com/gin-gonic/gin"
)

func NoRoute(c *gin.Context) {
	path := strings.Split(c.Request.URL.Path, "/")
	//fmt.Println(path)
	if (path[1] != "") && (path[1] == "api") {
		aRes := model.NewResponse()
		aRes.Code = http.StatusNotFound
		aRes.Msg = "not found"
		c.JSON(http.StatusNotFound, aRes)
	} else {
		c.HTML(http.StatusOK, "index.html", "")
	}
}
