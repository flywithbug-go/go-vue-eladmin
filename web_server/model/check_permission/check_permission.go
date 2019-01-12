package check_permission

import (
	"strings"
	"vue-admin/web_server/common"
	"vue-admin/web_server/model/model_user"

	"github.com/gin-gonic/gin"
)

func CheckPermission(c *gin.Context, permission string) bool {
	id := common.UserId(c)
	user := model_user.User{}
	user.Id = id
	user, err := user.FindOne()
	if err != nil {
		return false
	}
	for _, item := range user.RolesString {
		if strings.EqualFold(item, "") {

		}
		if strings.EqualFold(item, permission) {
			return true
		}
		if strings.HasSuffix(item, "ALL") {
			splits := strings.Split(item, "_")
			if strings.HasPrefix(item, splits[0]) {
				return true
			}
		}
	}
	return false
}
