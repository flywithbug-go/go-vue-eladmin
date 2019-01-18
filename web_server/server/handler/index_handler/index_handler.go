package index_handler

import (
	"time"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	time.Sleep(time.Second * 200)
}
