package model_monitor

import (
	"fmt"
	"time"
)

const (
	timeLayout = "2006-01-02"
)

func (l Log) AddMonitorInfo() {
	if len(l.UUID) == 0 {
		l.UUID = formatMd5(l.UserId, l.ClientIp)
	}
	l.Insert()
	if len(l.UUID) > 0 {
		visit := Visit{}
		visit.TimeDate = time.Now().Format(timeLayout)
		visit.UUID = l.UUID
		visit.ClientIp = l.ClientIp
		visit.IncrementVisit()
	}

}

func formatMd5(userId int64, clientIp string) string {
	return fmt.Sprintf("uid:%d-clientIp:%s", userId, clientIp)
}
