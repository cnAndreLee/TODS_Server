package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cnAndreLee/tods_server/common"
	"github.com/cnAndreLee/tods_server/model"
	"github.com/cnAndreLee/tods_server/response"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	db := common.GetDB()

	var RequestUser = model.User{}
	c.Bind(&RequestUser)
	fmt.Printf("收到注册请求，user: %v\n", RequestUser)
	RequestUser.Account = strings.Replace(RequestUser.Account, " ", "", -1)

	if RequestUser.Account == "" {
		response.Response(c, http.StatusBadRequest, 1, "user is null!", nil)
		return
	}

	if isAccountExist(RequestUser.Account) {
		fmt.Printf("用户已存在\n")
		response.Response(c, http.StatusBadRequest, 1, "user already exist", nil)
		return
	}

	//todo : 细化错误信息
	db.Debug().Create(&RequestUser)
	response.Response(c, http.StatusOK, 0, "ok", nil)

}

func JWTLogin(c *gin.Context) {

	var RequestUser = model.User{}
	c.Bind(&RequestUser)

	fmt.Printf("收到用户登录请求，accout:%v   key:%v\n", RequestUser.Account, RequestUser.Key)

	if strings.Replace(RequestUser.Account, " ", "", -1) == "" || strings.Replace(RequestUser.Key, " ", "", -1) == "" {
		response.Response(c, http.StatusBadRequest, 1, "account or password is null", nil)
		return
	}

	//todo : 校验用户
	var user model.User
	common.DB.Debug().Where("account = ?", RequestUser.Account).First(&user)
	if user.Key != RequestUser.Key {
		fmt.Printf("账号或密码错误\n")
		response.Response(c, http.StatusBadRequest, 1, "account or password is wrong", nil)
		return
	}

	//todo : outdate

	str, err := common.CreateToken(RequestUser.Account)
	if err != nil {
		response.ResponseServerError(c)
		return
	}

	//return token
	response.Response(c, http.StatusOK, 0, "ok", gin.H{"token": str})
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	response.Response(ctx, http.StatusOK, 0, "", gin.H{"user": user})
}

func isAccountExist(account string) bool {
	var user model.User
	if res := common.DB.Where("account = ?", account).First(&user); res.Error != nil {
		return false
	}

	return true
}
