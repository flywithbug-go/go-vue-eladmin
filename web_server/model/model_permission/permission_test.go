package model_permission

import (
	"encoding/json"
	"fmt"
	"testing"

	"vue-admin/web_server/core/mongo"
)

func TestPipe(t *testing.T) {
	mongo.DialMgo("127.0.0.1:27017")
	permission := Permission{}
	results, _ := permission.FindPipeAll()

	//result := make([]Permission, 0)
	//pipeline := []bson.M{
	//	{"$match": bson.M{"pid": 0}},
	//	{"$lookup": bson.M{"from": "permission", "localField": "_id", "foreignField": "pid", "as": "children"}},
	//}
	//
	//err := permission.pipeAll(pipeline, &result, true)
	//if err != nil {
	//	panic(err)
	//}
	js, _ := json.Marshal(results)
	fmt.Println(string(js))

	permission.Id = 10006
	permission, err := permission.FindOne()
	if err != nil {
		panic(err)
	}

	js, _ = json.Marshal(permission)
	fmt.Println(string(js))
}
