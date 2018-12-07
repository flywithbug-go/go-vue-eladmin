package model

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID		bson.ObjectId 	//`json:"id"`
	Name 	string  		//`json:"name"`
}
