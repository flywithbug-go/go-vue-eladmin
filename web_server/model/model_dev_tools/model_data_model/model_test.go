package model_data_model

import (
	"testing"
	"vue-admin/web_server/core/mongo"
)

func TestDataModel(t *testing.T) {
	mongo.RegisterMongo("127.0.0.1:27017", "doc_manager")

	dataModel := DataModel{}
	dataModel.Name = "User"
	dataModel.Desc = "test"
	//dataModel.Attributes = make([]Attributes, 0)

	a := Attribute{}
	a.Name = "Name"
	a.Type = modelAttributeTypeString

	dataModel.Attributes = append(dataModel.Attributes, a)
	_, err := dataModel.Insert()
	if err != nil {
		panic(err)
	}

}

func TestDataModel_Update(t *testing.T) {
	mongo.RegisterMongo("127.0.0.1:27017", "doc_manager")

	dataModel := DataModel{}
	dataModel.Name = "User"
	dataModel.Desc = "test"
	dataModel.Id = 10009

	a := Attribute{}
	a.Name = "Name1"
	a.Type = modelAttributeTypeString

	err := dataModel.AddAttribute(a)
	if err != nil {
		panic(err)
	}
}

func TestPipe(t *testing.T) {
	//mongo.RegisterMongo("127.0.0.1:27017", "doc_manager")
	//permission := model_permission.Permission{}
	//
	//list, err := permission.FetchTreeList(nil)
	//if err != nil {
	//	panic(err)
	//}
	//js, _ := json.Marshal(list)
	//fmt.Println(string(js))
}
