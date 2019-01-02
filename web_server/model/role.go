package model

const (
	roleCollection = "role"
)

//角色表，记录公司各种角色，比如：CEO 管理员，开发，开发经理，销售，销售主管，等
type Role struct {
	Id          int64  `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string `json:"name,omitempty"` //角色名称 CEO，CTO，主管，经理，程序员等。
	Code        string `json:"code"`           //角色编码
	DelFlag     bool   `json:"del_flag"`       //是否被删除
	Description string `json:"description"`
}
