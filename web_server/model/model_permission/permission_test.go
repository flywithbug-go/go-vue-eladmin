package model_permission

import (
	"encoding/json"
	"fmt"
	"testing"

	"gopkg.in/mgo.v2/bson"

	"vue-admin/web_server/core/mongo"
)

func TestPipe(t *testing.T) {
	mongo.DialMgo("127.0.0.1:27017")
	permission := Permission{}

	result := make([]Permission, 0)
	pipeline := []bson.M{
		{"$match": bson.M{"pid": 0}},
		{"$lookup": bson.M{"from": "permission", "localField": "_id", "foreignField": "pid", "as": "children"}},
	}

	err := permission.pipeAll(pipeline, &result, false)
	if err != nil {
		panic(err)
	}
	js, _ := json.Marshal(result)
	fmt.Println(string(js))

}
