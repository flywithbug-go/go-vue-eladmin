package log_writer

import (
	"fmt"
	"testing"
	"vue-admin/web_server/core/mongo"
)

func TestLog_Insert(t *testing.T) {
	mongo.RegisterMongo("127.0.0.1:27017", "log")

	log1 := new(Log)
	log1.Code = "23232"
	log1.Ext = Log{
		Code: "1010",
		Info: "name",
	}
	log := new(Log)
	log.Code = "23232"
	log.Ext = log1
	fmt.Println(log)
	log.Insert()
}
