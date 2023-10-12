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
	route.POST("/getUserRouter", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"data": gin.H{
				"router": []gin.H{{
					"id":        1,
					"name":      "index",
					"path":      "/index",
					"icon":      "HomeFilled",
					"component": "view/index/index",
					"meta": gin.H{
						"title": "首页",
					},
				}, {
					"id":        2,
					"name":      "list",
					"path":      "/list",
					"icon":      "StarFilled",
					"component": "view/list/index",
					"meta": gin.H{
						"title": "示例列表",
					},
				}, {
					"id":   3,
					"name": "system",
					"path": "/system",
					"icon": "Tools",
					"meta": gin.H{
						"title": "系统设置",
					},
					"children": []gin.H{{
						"id":        31,
						"name":      "systemRole",
						"path":      "role",
						"icon":      "Avatar",
						"component": "view/system/role",
						"meta": gin.H{
							"title": "角色管理",
						},
					}, {
						"id":        32,
						"name":      "systemUser",
						"path":      "user",
						"icon":      "UserFilled",
						"component": "view/system/user",
						"meta": gin.H{
							"title": "用户管理",
						},
					}, {
						"id":        33,
						"name":      "systemMenu",
						"path":      "menu",
						"icon":      "Menu",
						"component": "view/system/menu",
						"meta": gin.H{
							"title": "菜单管理",
						},
					}},
				}},
			},
			"msg": "获取成功",
		})
	})
}
