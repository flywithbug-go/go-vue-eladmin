package logger_handler

import (
	"fmt"
	"vue-admin/web_server/common"
	"vue-admin/web_server/log_writer"

	"github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
)

//log.Info("【GIN】【id:%d】【m:%s %s %s】【c:%s%3d%s】【l:%13v】【ip:%s】 【p:%s】【e:%s】【rid:%s】",

func Debug(c *gin.Context, format string, args ...interface{}) {
	xReqId := common.XRequestId(c)
	inf := ""
	if format != "" {
		inf = fmt.Sprintf(format, args...)
	} else {
		inf = fmt.Sprint(args...)
	}

	l := log_writer.DBlogPool.Get().(*log_writer.Log)
	l.RequestId = xReqId
	l.Info = inf
	log4go.DebugExt(l, fmt.Sprintf("【GIN】【reqID:%s】【info:%s】", xReqId, inf))
}

func Warn(c *gin.Context, format string, args ...interface{}) {
	xReqId := common.XRequestId(c)
	inf := ""
	if format != "" {
		inf = fmt.Sprintf(format, args...)
	} else {
		inf = fmt.Sprint(args...)
	}
	l := log_writer.DBlogPool.Get().(*log_writer.Log)
	l.RequestId = xReqId
	l.Info = inf
	log4go.WarnExt(l, fmt.Sprintf("【GIN】【reqID:%s】【info:%s】", xReqId, inf))
}

func Info(c *gin.Context, format string, args ...interface{}) {
	xReqId := common.XRequestId(c)
	inf := ""
	if format != "" {
		inf = fmt.Sprintf(format, args...)
	} else {
		inf = fmt.Sprint(args...)
	}
	l := log_writer.DBlogPool.Get().(*log_writer.Log)
	l.RequestId = xReqId
	l.Info = inf
	log4go.InfoExt(l, fmt.Sprintf("【GIN】【reqID:%s】【info:%s】", xReqId, inf))
}

func Error(c *gin.Context, format string, args ...interface{}) {
	xReqId := common.XRequestId(c)
	inf := ""
	if format != "" {
		inf = fmt.Sprintf(format, args...)
	} else {
		inf = fmt.Sprint(args...)
	}
	l := log_writer.DBlogPool.Get().(*log_writer.Log)
	l.RequestId = xReqId
	l.Info = inf
	log4go.ErrorExt(l, fmt.Sprintf("【GIN】【reqID:%s】【info:%s】", xReqId, inf))
}

func Fatal(c *gin.Context, format string, args ...interface{}) {
	xReqId := common.XRequestId(c)
	inf := ""
	if format != "" {
		inf = fmt.Sprintf(format, args...)
	} else {
		inf = fmt.Sprint(args...)
	}

	l := log_writer.DBlogPool.Get().(*log_writer.Log)
	l.RequestId = xReqId
	l.Info = inf
	log4go.FatalExt(l, fmt.Sprintf("【GIN】【reqID:%s】【info:%s】", xReqId, inf))
}
