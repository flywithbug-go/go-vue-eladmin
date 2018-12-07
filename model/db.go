package model

import (
	"github.com/flywithbug/log4go"
	"gopkg.in/mgo.v2"
)

var globalS *mgo.Session

func DialMgo(url string)  {
	s, err := mgo.Dial(url)
	if err != nil {
		log4go.Fatal("create session error ", err)
	}
	globalS = s
	log4go.Info("mongodb connected")
}
