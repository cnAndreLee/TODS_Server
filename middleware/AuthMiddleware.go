package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cnAndreLee/tods_server/common"
	"github.com/cnAndreLee/tods_server/model"
	"github.com/cnAndreLee/tods_server/response"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		fmt.Printf("token is %v\n", tokenString)

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response.Response(ctx, http.StatusUnauthorized, 1, "auth failed", nil)
		}

		tokenString = tokenString[7:]

		user, err := common.ParseToken(tokenString)
		if err != nil {
			response.Response(ctx, http.StatusUnauthorized, 1, "auth failed", nil)
			ctx.Abort()
			return
		}

		fmt.Printf("token parse success ,user is %v", user)
		// if !isAccountExist(user) {
		// 	response.Response(ctx, http.StatusUnauthorized, 1, "auth failed", nil)
		// 	ctx.Abort()
		// 	return
		// }

		ctx.Set("user", user)
		ctx.Next()
	}
}

func isAccountExist(account string) bool {
	var user model.User
	if res := common.DB.Where("account = ?", account).First(&user); res.Error != nil {
		return false
	}

	return true
}
