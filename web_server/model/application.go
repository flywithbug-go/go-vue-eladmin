package model

const (
	appCollection = "application"
)

type Application struct {
	Id 				int
	Name        	string  `json:"name"`	//应用（组件）名称
	Desc         	string //项目描述
	CreateTime   	int64  //创建时间
	Icon 			string	`json:"icon"` //icon 地址
	AppId 			string	`json:"app_id"`
	Owner 			string	`json:"owner"`  //负责人
}
