package user_handler

import (
	"net/http"
	"strconv"
	"strings"
	"vue-admin/web_server/common"
	"vue-admin/web_server/model"
	"vue-admin/web_server/model/model_user"
	"vue-admin/web_server/server/handler/check_permission"

	"github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

type UserRole struct {
	model_user.User
	Roles []string `json:"roles"`
}

type ParaUserEdit struct {
	model_user.User
	Enabled string `json:"enabled,omitempty" bson:"enabled,omitempty"`
}

func addUserHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()

	if check_permission.CheckNoPermission(c, model_user.UserPermissionCreate) {
		log4go.Info("has no permission")
		aRes.SetErrorInfo(http.StatusForbidden, "has no permission")
		return
	}
	user := new(model_user.User)
	err := c.BindJSON(user)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	err = user.Insert()
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	aRes.SetSuccess()
}

func getUserInfoHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	var id int64
	ids := c.Query("id")
	id, _ = strconv.ParseInt(ids, 10, 64)
	if id == 0 {
		id = common.UserId(c)
	} else {
		if check_permission.CheckNoPermission(c, model_user.UserPermissionSelect) {
			log4go.Info("has no permission")
			aRes.SetErrorInfo(http.StatusForbidden, "has no permission")
			return
		}
	}
	user := model_user.User{}
	user.Id = id
	user, err := user.FindOne()
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusUnauthorized, "user not found:"+err.Error())
		return
	}
	if !user.Enabled {
		log4go.Info("账号已停用")
		aRes.SetErrorInfo(http.StatusUnauthorized, "账号已停用，请联系管理员")
		return
	}
	roleUser := UserRole{}
	roleUser.User = user
	roleUser.Roles = user.RolesString
	if roleUser.Id == 10000 && len(roleUser.Roles) > 0 {
		roleUser.Roles = []string{"ADMIN"}
	}
	roleUser.RolesString = nil
	aRes.AddResponseInfo("user", roleUser)
}

func updateUserHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	if check_permission.CheckNoPermission(c, model_user.UserPermissionEdit) {
		log4go.Info("has no permission")
		aRes.SetErrorInfo(http.StatusForbidden, "has no permission")
		return
	}
	user := new(model_user.User)
	err := c.BindJSON(user)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	err = user.Update()
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "db update failed: "+err.Error())
		return
	}
	aRes.SetSuccessInfo(http.StatusOK, "success")
}

func deleteUserHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	if check_permission.CheckNoPermission(c, model_user.UserPermissionDelete) {
		log4go.Info("has no permission")
		aRes.SetErrorInfo(http.StatusForbidden, "has no permission")
		return
	}
	user := new(model_user.User)
	err := c.BindJSON(user)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	if common.UserId(c) == user.Id {
		log4go.Info("can not delete your self")
		aRes.SetErrorInfo(http.StatusForbidden, "can not delete your self")
		return
	}
	err = user.Remove()
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "db delete failed: "+err.Error())
		return
	}
	aRes.SetSuccessInfo(http.StatusOK, "success")
}

func getUserListInfoHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()

	if check_permission.CheckNoPermission(c, model_user.UserPermissionSelect) {
		log4go.Info("has no permission")
		aRes.SetErrorInfo(http.StatusForbidden, "has no permission")
		return
	}
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	sort := c.Query("sort")
	name := c.Query("username")
	email := c.Query("email")
	enabled := c.Query("enabled")

	if strings.EqualFold(sort, "-id") {
		sort = "-_id"
	} else if strings.EqualFold(sort, "+id") {
		sort = "+_id"
	} else if len(sort) == 0 {
		sort = "+_id"
	}
	if limit == 0 {
		limit = 10
	}
	if page != 0 {
		page--
	}
	userId := common.UserId(c)
	if userId <= 0 {
		aRes.SetErrorInfo(http.StatusUnauthorized, "user not found")
		return
	}
	query := bson.M{}
	if len(name) > 0 {
		query["username"] = bson.M{"$regex": name, "$options": "i"}
	}
	if len(email) > 0 {
		query["email"] = bson.M{"$regex": email, "$options": "i"}
	}

	if strings.EqualFold(enabled, "true") {
		query["enabled"] = true
	}
	if strings.EqualFold(enabled, "false") {
		query["enabled"] = false
	}

	var user = model_user.User{}
	totalCount, _ := user.TotalCount(query, nil)
	appList, err := user.FindPageFilter(page, limit, query, nil, sort)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, err.Error())
		return
	}
	aRes.AddResponseInfo("list", appList)
	aRes.AddResponseInfo("total", totalCount)
}
