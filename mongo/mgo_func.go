package mongo

import (
	"errors"
	"gopkg.in/mgo.v2"
)

func DialMongo(url ,db string) (*mgo.Session,error)  {
	if db == "" {
		return nil, errors.New("db name can not be nil")
	}
	return mgo.Dial(url)
}
