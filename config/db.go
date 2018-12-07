package config

import (
	"doc-manager/mongo"
	"github.com/flywithbug/log4go"
)
var varMongoInited = false
func MongoInited() bool {
	return varMongoInited
}
func DailMongo()  {
  	err :=	mongo.DialMongo("default","127.0.0.1:27017","doc-manager")
	if err != nil {
		log4go.Error(err.Error())
		panic(err)
	}
  	varMongoInited = true
  	log4go.Info("mongodb connected")
}
