package response

import (
	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpStatus int, status int, msg string, data gin.H) {
	ctx.JSON(httpStatus, gin.H{"status": status, "msg": msg, "data": data})
}

func ResponseServerError(ctx *gin.Context) {
	Response(ctx, 500, 1, "server error", nil)

}

// func Success(ctx *gin.Context, code int, data gin.H, msg string) {
// 	Response(ctx, http.StatusOK, 200, data, msg)
// }

// func Fail(ctx *gin.Context, code int, data gin.H, msg string) {
// 	Response(ctx, http.StatusOK, 400, data, msg)
// }
