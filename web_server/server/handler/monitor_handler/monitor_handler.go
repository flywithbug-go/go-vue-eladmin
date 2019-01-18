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
	timeLayout = "2006-01-02"
)

type responseVisit struct {
	DayVisit   int `json:"day_visit,omitempty"`
	TotalVisit int `json:"total_visit,omitempty"`
	DayApi     int `json:"day_api,omitempty"`
	TotalApi   int `json:"total_api,omitempty"`

	DayIP   int `json:"day_ip,omitempty"`
	TotalIp int `json:"total_ip,omitempty"`
}

func visitHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponseCode, aRes.Code)
		c.JSON(http.StatusOK, aRes)
	}()
	query := bson.M{}
	resVisit := responseVisit{}
	timeF := time.Now().Format(timeLayout)
	vApi := model_monitor.VisitApi{}
	vUId := model_monitor.VisitUId{}

	query = bson.M{"time_date": bson.M{"$regex": timeF, "$options": "i"}}
	resVisit.DayIP, _ = vApi.TotalSumCount(query)

	resVisit.DayVisit, _ = vUId.TotalSumCount(query)
	resVisit.DayIP, _ = vUId.TotalCount(query, nil)
	query = bson.M{"time_date": bson.M{"$regex": "", "$options": "i"}}
	resVisit.TotalVisit, _ = vUId.TotalSumCount(query)
	resVisit.TotalIp, _ = vUId.TotalCount(nil, nil)
	resVisit.TotalApi, _ = vApi.TotalCount(nil, nil)
	aRes.AddResponseInfo("visit", resVisit)
}
