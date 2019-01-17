package handler_common

import "github.com/gin-gonic/gin"

type StateType int

const (
	RouterTypeNormal StateType = iota
	RouterTypeNeedAuth
)

type GinHandleFunc struct {
	Handler    gin.HandlerFunc
	RouterType StateType
	Method     string
	Route      string
}
