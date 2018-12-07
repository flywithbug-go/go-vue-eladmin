package mongo

import (
	"errors"
	"gopkg.in/mgo.v2"
)

 const DefaultSessionName = "default"

func DialMongo(alias,url ,db string) (error)  {
	if _, ok := sessionMap[DefaultSessionName]; !ok && alias != DefaultSessionName {
		return errors.New(ErrNoDefaultConnection)
	}
	if _, ok := sessionMap[alias]; ok {
		return errors.New(ErrExistConnectionAlias)
	}
	if db == "" {
		return errors.New("db name can not be nil")
	}
	aSession , err := mgo.Dial(url)
	if err != nil {
		return err
	}
	aMongo := new(sessionMgo)
	aMongo.session = aSession
	aMongo.db = db
	sessionMap[alias] = aMongo
	return nil
}

func NewMongo(alias string)Mongoer  {
	aMongo := new(sessionMgo)
	if alias == "" {
		 alias = DefaultSessionName
	}
	aMongo.Use(alias)
	return aMongo
}
