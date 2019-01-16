package logger_handler

import (
	"fmt"
	"vue-admin/web_server/common"

	"github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
)

func Debug(c *gin.Context, format string, args ...interface{}) {
	xReqId := common.XRequestId(c)
	inf := ""
	if format != "" {
		inf = fmt.Sprintf(format, args...)
	} else {
		inf = fmt.Sprint(args...)
	}
	log4go.Info(fmt.Sprintf("【GIN】\t【reqID:%s】\t【info:%s】", xReqId, inf))
}

func Warn(c *gin.Context, format string, args ...interface{}) {
	xReqId := common.XRequestId(c)
	inf := ""
	if format != "" {
		inf = fmt.Sprintf(format, args...)
	} else {
		inf = fmt.Sprint(args...)
	}
	log4go.Info(fmt.Sprintf("【GIN】\t【reqID:%s】\t【info:%s】", xReqId, inf))
}

func Info(c *gin.Context, format string, args ...interface{}) {
	xReqId := common.XRequestId(c)
	inf := ""
	if format != "" {
		inf = fmt.Sprintf(format, args...)
	} else {
		inf = fmt.Sprint(args...)
	}
	log4go.Info(fmt.Sprintf("【GIN】\t【reqID:%s】\t【info:%s】", xReqId, inf))
}

func Error(c *gin.Context, format string, args ...interface{}) {
	xReqId := common.XRequestId(c)
	inf := ""
	if format != "" {
		inf = fmt.Sprintf(format, args...)
	} else {
		inf = fmt.Sprint(args...)
	}
	log4go.Info(fmt.Sprintf("【GIN】\t【reqID:%s】\t【info:%s】", xReqId, inf))
}

func Fatal(c *gin.Context, format string, args ...interface{}) {
	xReqId := common.XRequestId(c)
	inf := ""
	if format != "" {
		inf = fmt.Sprintf(format, args...)
	} else {
		inf = fmt.Sprint(args...)
	}
	log4go.Info(fmt.Sprintf("【GIN】\t【reqID:%s】\t【info:%s】", xReqId, inf))
}
