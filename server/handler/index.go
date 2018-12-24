package handler

import (
	"doc-manager/model"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


func NoRoute(c *gin.Context) {
	path := strings.Split(c.Request.URL.Path, "/")
	fmt.Println(path)
	if (path[1] != "") && (path[1] == "api") {
		aRes := model.NewResponse()
		aRes.Code = http.StatusNotFound
		aRes.Msg = "no route"
		c.JSON(http.StatusNotFound, aRes)
	} else {
		c.HTML(http.StatusOK, "index.html", "")
	}
}