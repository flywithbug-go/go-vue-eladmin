package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func htmlHandler(c *gin.Context) {
	//path := strings.Split(c.Request.URL.Path, "/")
	//fmt.Println(path)
	c.HTML(http.StatusOK, "index.html", "")

}