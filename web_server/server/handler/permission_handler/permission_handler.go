package permission_handler

import (
	"net/http"
	"strconv"
	"vue-admin/web_server/model"
	"vue-admin/web_server/model/model_permission"
	"vue-admin/web_server/model/mongo_index"

	"github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func addPermissionHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	p := model_permission.Permission{}
	err := c.BindJSON(&p)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}

	err = p.Insert()
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, "server invalid: "+err.Error())
		return
	}
}

func getPermissionHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	ids := c.Query("id")
	var mInfo = model_permission.Permission{}
	id, _ := strconv.Atoi(ids)
	mInfo.Id = int64(id)
	mInfo, err := mInfo.FindOne()
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "invalid: "+err.Error())
		return
	}
	aRes.AddResponseInfo("permission", mInfo)
}

func updatePermissionHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	para := model_permission.Permission{}
	err := c.BindJSON(&para)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	err = para.Update()
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "invalid: "+err.Error())
		return
	}
}

func removePermissionHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	//need id
	para := model_permission.Permission{}
	err := c.BindJSON(&para)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	err = para.Remove()
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "invalid: "+err.Error())
		return
	}
}

func getPermissionListHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	var per = model_permission.Permission{}

	//name := c.Query("name")
	//sort := bson.M{"$sort": bson.M{"_id": 1}}
	//match := bson.M{"$match": bson.M{"pid": 0}}
	//if len(name) > 0 {
	//	match = bson.M{"$match": bson.M{"pid": 0, "name": bson.M{"$regex": name, "$options": "i"}}}
	//}
	//lookup := bson.M{"$graphLookup": bson.M{"from": mongo_index.CollectionPermission, "startWith": "$_id", "connectFromField": "_id", "connectToField": "pid", "as": "children"}}
	//pipeline := []bson.M{
	//	match,
	//	sort,
	//	lookup,
	//}
	results, err := per.FindAllList()
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusUnauthorized, "app version list find error"+err.Error())
		return
	}
	aRes.AddResponseInfo("list", results)
}

func getPermissionTreeHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	name := c.Query("name")
	sort := bson.M{"$sort": bson.M{"_id": 1}}
	match := bson.M{"$match": bson.M{"pid": 0}}
	if len(name) > 0 {
		match = bson.M{"$match": bson.M{"pid": 0, "name": bson.M{"$regex": name, "$options": "i"}}}
	}
	lookup := bson.M{"$lookup": bson.M{"from": mongo_index.CollectionPermission, "localField": "_id", "foreignField": "pid", "as": "children"}}
	var per = model_permission.Permission{}
	pipeline := []bson.M{
		match,
		sort,
		lookup,
	}
	results, err := per.FindPipeline(pipeline)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusUnauthorized, "app version list find error"+err.Error())
		return
	}
	aRes.AddResponseInfo("list", results)
}
