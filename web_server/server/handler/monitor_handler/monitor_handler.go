package monitor_handler

import (
	"net/http"
	"time"
	"vue-admin/web_server/common"
	"vue-admin/web_server/model"
	"vue-admin/web_server/model/model_monitor"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
)

const (
	timeLayout   = "2006-01-02"
	MonitorVisit = "visit"
)

type responseVisit struct {
	DayVisit   int `json:"dayVisit"`
	TotalVisit int `json:"totalVisit"`
	DayApi     int `json:"dayApi"`
	TotalApi   int `json:"totalApi"`

	DayIP   int `json:"dayIp"`
	TotalIp int `json:"totalIp"`
}

func visitHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponseCode, aRes.Code)
		c.JSON(http.StatusOK, aRes)
	}()
	query := bson.M{}
	monitorCount := model_monitor.MonitorCount{}

	resVisit := responseVisit{}
	timeF := time.Now().Format(timeLayout)
	//log4go.Info(timeF)
	vApi := model_monitor.VisitApi{}
	vUId := model_monitor.VisitUId{}

	query = bson.M{"time_date": bson.M{"$regex": timeF, "$options": "i"}}
	resVisit.DayApi, _ = vApi.TotalSumCount(query)
	resVisit.DayVisit, _ = monitorCount.TotalSumCount(query) //日访问
	resVisit.DayIP, _ = vUId.TotalCount(query, nil)

	query = bson.M{"time_date": bson.M{"$regex": "", "$options": "i"}}
	resVisit.TotalApi, _ = vApi.TotalSumCount(query)
	resVisit.TotalVisit, _ = monitorCount.TotalSumCount(query) //总访问

	resVisit.TotalIp, _ = vUId.TotalCount(nil, nil)

	aRes.AddResponseInfo("visit", resVisit)
}

func visitCountHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponseCode, aRes.Code)
		c.JSON(http.StatusOK, aRes)
	}()

	mon := model_monitor.MonitorCount{}
	mon.Monitor = MonitorVisit
	mon.IncrementMonitorCount()
	aRes.SetSuccess()
}

func chartListHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponseCode, aRes.Code)
		c.JSON(http.StatusOK, aRes)
	}()

}
