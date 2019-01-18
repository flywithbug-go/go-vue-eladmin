package model_monitor

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	timeLayout = "2006-01-02 15:04:05"
)

var (
	visitApiPool *sync.Pool
	visitUIdPool *sync.Pool
)

func init() {
	visitApiPool = &sync.Pool{New: func() interface{} {
		return &VisitApi{}
	}}
	visitUIdPool = &sync.Pool{New: func() interface{} {
		return &VisitUId{}
	}}
}

func (l Log) AddMonitorInfo() {
	visitUid := visitUIdPool.Get().(*VisitUId)
	vApi := visitApiPool.Get().(*VisitApi)
	if len(l.UUID) == 0 {
		visitUid.findOne(bson.M{"user_id": l.UserId}, nil)
		l.UUID = visitUid.UUID
	}
	l.Insert()
	timeF := time.Now().Format(timeLayout)
	if l.Latency > 0 {
		if len(l.UUID) > 0 {
			visitUid.TimeDate = timeF[:10]
			visitUid.UUID = l.UUID
			visitUid.UserId = l.UserId
			visitUid.ClientIp = l.ClientIp
			visitUid.IncrementVisitUId()
		}
		if len(l.Path) > 0 {
			vApi.TimeDate = timeF[:13]
			vApi.Method = l.Method
			list := strings.Split(l.Path, "?")
			vApi.Path = list[0]
			vApi.IncrementVisitApi()
		}
	}
	visitUIdPool.Put(visitUid)
	visitApiPool.Put(vApi)
}

func formatMd5(userId int64, clientIp string) string {
	return fmt.Sprintf("uid:%d-clientIp:%s", userId, clientIp)
}
