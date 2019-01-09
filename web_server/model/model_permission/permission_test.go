package model_permission

import (
	"encoding/json"
	"fmt"
	"testing"

	"gopkg.in/mgo.v2/bson"

	"vue-admin/web_server/core/mongo"
)

type PermissionResult struct {
	Id       int64        `json:"_id,omitempty" bson:"_id,omitempty"`
	PId      int64        `json:"pid,omitempty" bson:"pid,omitempty"`
	Name     string       `json:"name"`
	Alias    string       `json:"alias"`
	Children []Permission `json:"children"`
}

func TestPipe(t *testing.T) {
	mongo.DialMgo("127.0.0.1:27017")
	permission := Permission{}

	result := make([]PermissionResult, 0)
	pipeline := []bson.M{
		{"$match": bson.M{"name": "user_all"}},
		{"$lookup": bson.M{"from": "permission", "localField": "_id", "foreignField": "pid", "as": "children"}},
	}

	err := permission.pipeAll(pipeline, &result, false)
	if err != nil {
		panic(err)
	}
	js, _ := json.Marshal(result)
	fmt.Println(string(js))

}
