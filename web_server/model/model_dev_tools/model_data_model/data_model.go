package model_data_model

type typeStatus int

const (
	modelAttributeTypeUndefine typeStatus = iota //待定
	//基础类型
	modelAttributeTypeInt    //Int类型
	modelAttributeTypeBool   //布尔类型
	modelAttributeTypeString //String类型
	modelAttributeTypeList   //数组 （基础类型或者模型）

	modelAttributeTypeObject //模型

)

type DataModel struct {
	Id         int64       `json:"id,omitempty" bson:"_id,omitempty"`
	Desc       string      `json:"desc"`
	Name       string      `json:"name"`
	CreateTime int64       `json:"create_time"`
	AppIds     []int64     `json:"app_ids"` //所属应用
	Attributes []Attribute //模型的属性表

}

type list struct {
	ModelId   string
	ModelName string
}

type Attribute struct {
	Type      int         `json:"type"` //int string list
	Name      string      `json:"name"`
	ModelName string      `json:"model_name"`
	Default   interface{} `json:"default"`
}
