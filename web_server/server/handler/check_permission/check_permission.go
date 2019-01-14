package check_permission

import (
	"strings"
	"vue-admin/web_server/common"
	"vue-admin/web_server/model/model_user"

	"github.com/gin-gonic/gin"
)

const (
	SUPERADMIN = "ADMIN"
)

func CheckNoPermission(c *gin.Context, permission string) bool {
	id := common.UserId(c)
	user := model_user.User{}
	user.Id = id
	user, err := user.FindOne()
	if err != nil {
		return true
	}
	for index := range user.RolesString {
		item := user.RolesString[index]
		if strings.EqualFold(item, SUPERADMIN) {
			return false
		}
		if strings.EqualFold(item, permission) {
			return false
		}
		if strings.HasSuffix(item, "ALL") {
			splits := strings.Split(item, "_")
			if strings.HasPrefix(item, splits[0]) {
				return false
			}
		}
	}
	return true
}
