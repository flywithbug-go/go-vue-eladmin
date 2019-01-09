package model_menu

type Menu struct {
	Name       string
	Path       string
	Redirect   string
	Component  string
	AlwaysShow bool `json:"alwaysShow"`
	Meta       Meta `json:"meta"`
}
