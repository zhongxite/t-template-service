package router

import (
	"github.com/gin-gonic/gin"
	"zhongxite/t-template/middleware"
	"zhongxite/t-template/service"
)

func MenusRouter(r *gin.RouterGroup) {
	route := r.Group("/menu")
	route.POST("/getMenusList", middleware.Auth(), service.GetMenusList)
	route.POST("/menusAddOrModify", middleware.Auth(), service.MenusAddOrModify)
	route.POST("/deleteMenus", middleware.Auth(), service.DeleteMenus)
}
