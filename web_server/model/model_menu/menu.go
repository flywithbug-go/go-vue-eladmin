package model_menu

type Menu struct {
	Id         int64
	PId        int64 //父节点ID
	Sort       int
	Name       string
	Path       string
	Redirect   string
	Component  string
	AlwaysShow bool `json:"alwaysShow"`
	Meta       Meta `json:"meta"`
}
