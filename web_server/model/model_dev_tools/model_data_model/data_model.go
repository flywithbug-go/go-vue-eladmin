package model_data_model

type DataModel struct {
	Id         int64  `json:"id,omitempty" bson:"_id,omitempty"`
	Desc       string `json:"desc"`
	Name       string `json:"name"`
	CreateTime int64  `json:"create_time"`
}
