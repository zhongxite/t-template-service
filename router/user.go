package router

import (
	"github.com/gin-gonic/gin"
	"zhongxite/t-template/middleware"
	"zhongxite/t-template/service"
)

func UserRouter(r *gin.RouterGroup) {
	route := r.Group("/user")
	route.POST("/createUser", service.CreateUser)
	route.POST("/login", service.Login)
	route.POST("/getUserInfo", middleware.Auth(), service.GetUserInfo)
	route.POST("/updateUserInfo", middleware.Auth(), service.UpdateUserInfo)
	route.GET("/getUserList", middleware.Auth(), service.GetUserList)
	route.POST("/deleteUser", middleware.Auth(), service.DeleteUser)
}
