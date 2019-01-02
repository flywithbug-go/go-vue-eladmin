package model

//权限表 type
//1. Delete ReadWrite
//2. Read
//3. NoRight

type Permission struct {
	Id          int64  `json:"id,omitempty" bson:"_id,omitempty"`
	Type        int    `json:"type"`           //
	Name        string `json:"name,omitempty"` //
	Code        int    `json:"code"`           //
	DelFlag     bool   `json:"del_flag"`       //
	Description string `json:"description"`
}
