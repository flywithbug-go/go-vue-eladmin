package middleware

import (
	"encoding/base64"
	"encoding/binary"
	"time"
	"vue-admin/web_server/common"

	log "github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
)

type RequestHeader struct {
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		xReqId := common.XRequestId(c)
		if xReqId == "" {
			xReqId = GenReqID()
		}
		c.Set(common.KeyContextRequestId, xReqId)
		//headers, _ := json.Marshal(c.Request.Header)
		//log.Info("[GIN] [%s] [Started]\tRequestHeader::%s\n", xReqId, headers)
		// Process request
		c.Next()
		end := time.Now()
		latency := end.Sub(start)

		statusCode := c.Writer.Status()
		statusColor := colorForStatus(statusCode)
		clientIP := c.ClientIP()
		method := c.Request.Method
		methodColor := colorForMethod(method)
		comment := c.Errors.ByType(gin.ErrorTypePrivate).String()
		userId := common.UserId(c)
		log.Info("【GIN】\t【id:%d】\t【reqID:%s】 【method:%s %s %s】\t 【code:%s%3d%s】\t【latency:%13v】\t【IP:%s】 【path:%s】\t【gError:%s】",
			userId,
			xReqId,
			methodColor, method, reset,
			statusColor, statusCode, reset,
			latency,
			clientIP,
			c.Request.URL.String(),
			comment,
		)
	}
}

var pid = uint32(time.Now().UnixNano() % 4294967291)

// GenReqID is a random generate string func
func GenReqID() string {
	var b [12]byte
	binary.LittleEndian.PutUint32(b[:], pid)
	binary.LittleEndian.PutUint64(b[4:], uint64(time.Now().UnixNano()))
	return base64.URLEncoding.EncodeToString(b[:])
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
