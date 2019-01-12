package user_handler

import (
	"net/http"
	"strconv"
	"strings"
	"vue-admin/web_server/common"
	"vue-admin/web_server/core/jwt"
	"vue-admin/web_server/model"
	"vue-admin/web_server/model/check_permission"
	"vue-admin/web_server/model/model_user"
	"vue-admin/web_server/server/sync"

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

func loginHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	user := model_user.User{}
	err := c.BindJSON(&user)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	user, err = model_user.LoginUser(user.Username, user.Password)
	if err != nil {
		log4go.Error(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "username or password not right")
		return
	}
	claims := jwt.NewCustomClaims(user.Id, user.Username)
	token, err := jwt.GenerateToken(claims)
	if err != nil {
		log4go.Error(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "token generate error"+err.Error())
		return
	}
	userAg := c.GetHeader(common.KeyUserAgent)
	_, err = model_user.UserLogin(user.Id, userAg, token, c.ClientIP())
	if err != nil {
		log4go.Error(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "token generate error"+err.Error())
		return
	}
	sync.SetKeyValue(token)
	aRes.SetResponseDataInfo("token", token)
}

func registerHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	user := new(model_user.User)
	err := c.BindJSON(user)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	if user.Username == "" {
		aRes.SetErrorInfo(http.StatusBadRequest, "username can not be nil")
		return
	}
	if user.Password == "" {
		aRes.SetErrorInfo(http.StatusBadRequest, "Password can not be nil")
		return
	}
	if user.Email == "" {
		aRes.SetErrorInfo(http.StatusBadRequest, "email can not be nil")
		return
	}
	err = user.Insert()
	if err != nil {
		aRes.SetErrorInfo(http.StatusBadRequest, err.Error())
		return
	}
	aRes.AddResponseInfo("user", user)
}

func logoutHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	token := common.UserToken(c)
	if token == "" {
		log4go.Info("token not found")
		aRes.SetErrorInfo(http.StatusBadRequest, "token not found")
		return
	}
	sync.RemoveKey(token)
	err := model_user.UpdateLoginStatus(token, model_user.StatusLogout)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, err.Error())
		return
	}
	aRes.SetSuccessInfo(http.StatusOK, "success")
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
	name := c.Query("name")
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
		query["name"] = bson.M{"$regex": name, "$options": "i"}
	}
	var user = model_user.User{}
	totalCount, _ := user.TotalCount(query, nil)
	appList, err := user.FindPageFilter(page, limit, query, nil, sort)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusUnauthorized, "apps find error"+err.Error())
		return
	}
	aRes.AddResponseInfo("list", appList)
	aRes.AddResponseInfo("total", totalCount)
}
