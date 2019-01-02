package model

const (
	userRoleCollection = "user_role"
)

type UserRole struct {
	Id     int64 `json:"id" bson:"_id"`
	UserId int64
	RoleId int64
}
