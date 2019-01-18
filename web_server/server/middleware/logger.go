package middleware

import (
	"fmt"
	"time"
	"vue-admin/web_server/common"
	"vue-admin/web_server/log_writer"

	"gopkg.in/mgo.v2/bson"

	log "github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
)

//对照表 	rid		id		m		c 			l		p
// 		xReqId  userId 	method 	statusCode	latency	path
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		l := log_writer.DBlogPool.Get().(*log_writer.Log)
		// Start timer
		start := time.Now()
		l.RequestId = c.Request.Header.Get(common.KeyContextRequestId)
		if l.RequestId == "" {
			l.RequestId = GenReqID()
		}
		l.UUID = c.Request.Header.Get(common.KeyUUID)
		c.Header(common.KeyContextRequestId, l.RequestId)
		c.Set(common.KeyContextRequestId, l.RequestId)
		l.StartTime = start.UnixNano()
		l.ClientIp = c.ClientIP()
		l.Method = c.Request.Method
		l.Path = c.Request.URL.String()

		methodColor := colorForMethod(l.Method)
		log.InfoExt(l, "【GIN】【Start】【rid:%s】【m:%s %s %s】【ip:%s】 【p:%s】",
			l.RequestId,
			methodColor, l.Method, reset,
			l.ClientIp,
			l.Path)
		//----====----
		c.Next()
		end := time.Now()
		l.EndTime = end.UnixNano()
		l.Latency = end.Sub(start)
		l.StatusCode = c.Writer.Status()
		statusColor := colorForStatus(l.StatusCode)
		comment := c.Errors.ByType(gin.ErrorTypePrivate).String()
		l.UserId = common.UserId(c)
		l.Info = fmt.Sprintf("【GIN】【Completed】【id:%d】【rid:%s】【m:%s】【c:%3d】【l:%13v】【ip:%s】 【p:%s】【e:%s】",
			l.UserId,
			l.RequestId,
			l.Method,
			l.StatusCode,
			l.Latency,
			l.ClientIp,
			l.Path,
			comment)

		l.Para = common.Para(c)
		l.ResponseCode = common.ResponseCode(c)
		log.InfoExt(l, "【GIN】【Completed】【id:%d】【rid:%s】【m:%s %s %s】【c:%s%3d%s】【l:%13v】【ip:%s】 【p:%s】【e:%s】",
			l.UserId,
			l.RequestId,
			methodColor, l.Method, reset,
			statusColor, l.StatusCode, reset,
			l.Latency,
			l.ClientIp,
			l.Path,
			comment,
		)
	}
}

// GenReqID is a random generate string func
func GenReqID() string {
	return bson.NewObjectId().String()
}

var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

// ErrorLogger func
func ErrorLogger() gin.HandlerFunc {
	return ErrorLoggerT(gin.ErrorTypeAny)
}

// ErrorLoggerT func
func ErrorLoggerT(typ gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		errors := c.Errors.ByType(typ)
		if len(errors) > 0 {
			c.JSON(-1, errors)
		}
	}
}

func colorForStatus(code int) string {
	switch {
	case code >= 200 && code < 300:
		return green
	case code >= 300 && code < 400:
		return white
	case code >= 400 && code < 500:
		return yellow
	default:
		return red
	}
}

func colorForMethod(method string) string {
	switch method {
	case "GET":
		return blue
	case "POST":
		return cyan
	case "PUT":
		return yellow
	case "DELETE":
		return red
	case "PATCH":
		return green
	case "HEAD":
		return magenta
	case "OPTIONS":
		return white
	default:
		return reset
	}
}

var (
	skipPaths = []string{""}
)
