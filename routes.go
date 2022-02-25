package main

import (
	"github.com/cnAndreLee/tods_server/controller"
	"github.com/cnAndreLee/tods_server/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	r.Use(middleware.CORSMiddleware())
	r.Handle("POST", "/api/token", controller.JWTLogin)
	r.Handle("POST", "/api/register", controller.Register)
	r.Handle("GET", "/api/info", middleware.AuthMiddleware(), controller.Info)

	return r

}
