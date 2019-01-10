package user_handler

import (
	"net/http"
	"strconv"
	"strings"
	"vue-admin/web_server/common"
	"vue-admin/web_server/core/jwt"
	"vue-admin/web_server/model"
	"vue-admin/web_server/server/sync"

	"gopkg.in/mgo.v2/bson"

	"github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
)

type UserRole struct {
	model.User
	Roles []string `json:"roles"`
}

func loginHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	user := model.User{}
	err := c.BindJSON(&user)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	user, err = model.LoginUser(user.Account, user.Password)
	if err != nil {
		log4go.Error(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "account or password not right")
		return
	}
	claims := jwt.NewCustomClaims(user.Id, user.Account)
	token, err := jwt.GenerateToken(claims)
	if err != nil {
		log4go.Error(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "token generate error"+err.Error())
		return
	}
	userAg := c.GetHeader(common.KeyUserAgent)
	_, err = model.UserLogin(user.Id, userAg, token, c.ClientIP())
	if err != nil {
		log4go.Error(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "token generate error"+err.Error())
		return
	}
	sync.SetKeyValue(token)
	aRes.SetResponseDataInfo("token", token)
	aRes.AddResponseInfo("user", user)

}

func registerHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	user := new(model.User)
	err := c.BindJSON(user)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	if user.Account == "" {
		aRes.SetErrorInfo(http.StatusBadRequest, "account can not be nil")
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
	err := model.UpdateLoginStatus(token, model.StatusLogout)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, err.Error())
		return
	}
	aRes.SetSuccessInfo(http.StatusOK, "success")
}

func getUserInfoHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	userId := common.UserId(c)
	if userId <= 0 {
		log4go.Info("user not found")
		aRes.SetErrorInfo(http.StatusUnauthorized, "user not found")
		return
	}
	user, err := model.FindByUserId(userId)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusUnauthorized, "user not found:"+err.Error())
		return
	}
	roleUser := UserRole{}
	roleUser.User = user
	roleUser.Roles = []string{"admin"}
	aRes.AddResponseInfo("user", roleUser)
}

func getUserListInfoHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()

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
	var user = model.User{}
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

func updateUserHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.JSON(http.StatusOK, aRes)
	}()
	user := new(model.User)
	err := c.BindJSON(user)
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "para invalid"+err.Error())
		return
	}
	user.Id = common.UserId(c)
	err = user.Update()
	if err != nil {
		log4go.Info(err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, "db update failed: "+err.Error())
		return
	}
	aRes.SetSuccessInfo(http.StatusOK, "success")
}

func searchUserHandler(c *gin.Context) {

}
