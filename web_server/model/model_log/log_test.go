package model_log

import (
	"fmt"
	"testing"
	"vue-admin/web_server/core/mongo"
)

func TestLog_Insert(t *testing.T) {
	mongo.RegisterMongo("127.0.0.1:27017", "log")

	log := new(Log)
	log.Code = "23232"
	fmt.Println(log)
	log.Insert()
}
